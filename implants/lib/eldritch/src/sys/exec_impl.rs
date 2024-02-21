use super::super::insert_dict_kv;
use super::CommandOutput;
use anyhow::{Context, Result};
use starlark::{
    collections::SmallMap,
    const_frozen_string,
    values::{dict::Dict, Heap},
};
use std::process::Command;
#[cfg(any(target_os = "linux", target_os = "macos", target_os = "freebsd"))]
use {
    nix::unistd::setsid,
    nix::unistd::{fork, ForkResult},
    std::process::{exit, Stdio},
};

pub fn exec(
    starlark_heap: &Heap,
    path: String,
    args: Vec<String>,
    disown: Option<bool>,
) -> Result<Dict> {
    let cmd_res = handle_exec(path, args, disown)?;

    let res = SmallMap::new();
    let mut dict_res = Dict::new(res);
    insert_dict_kv!(dict_res, starlark_heap, "stdout", cmd_res.stdout, String);
    insert_dict_kv!(dict_res, starlark_heap, "stderr", cmd_res.stderr, String);
    insert_dict_kv!(dict_res, starlark_heap, "stauts", cmd_res.status, i32);

    Ok(dict_res)
}

fn handle_exec(path: String, args: Vec<String>, disown: Option<bool>) -> Result<CommandOutput> {
    let should_disown = disown.unwrap_or(false);

    if !should_disown {
        let res = Command::new(path).args(args).output()?;

        let res = CommandOutput {
            stdout: String::from_utf8(res.stdout)?,
            stderr: String::from_utf8(res.stderr)?,
            status: res
                .status
                .code()
                .context("Failed to retrieve status code")?,
        };
        Ok(res)
    } else {
        #[cfg(target_os = "windows")]
        return Err(anyhow::anyhow!(
            "Windows is not supported for disowned processes."
        ));

        #[cfg(any(target_os = "linux", target_os = "macos", target_os = "freebsd"))]
        match unsafe { fork()? } {
            ForkResult::Parent { child } => {
                if child.as_raw() < 0 {
                    return Err(anyhow::anyhow!("Pid was negative. ERR".to_string()));
                }
                Ok(CommandOutput {
                    stdout: "".to_string(),
                    stderr: "".to_string(),
                    status: 0,
                })
            }
            ForkResult::Child => {
                setsid()?;
                match unsafe { fork()? } {
                    ForkResult::Parent { child } => {
                        if child.as_raw() < 0 {
                            return Err(anyhow::anyhow!("Pid was negative. ERR".to_string()));
                        }
                        exit(0);
                    }
                    ForkResult::Child => {
                        let _res = Command::new(path)
                            .args(args)
                            .stdin(Stdio::null())
                            .stdout(Stdio::null())
                            .stderr(Stdio::null())
                            .spawn()?;
                        exit(0);
                    }
                }
            }
        }
    }
}

#[cfg(test)]
mod tests {
    use std::{fs, path::Path, thread, time};

    use tempfile::NamedTempFile;

    use super::*;
    #[test]
    fn test_sys_exec_current_user() -> anyhow::Result<()> {
        if cfg!(target_os = "linux")
            || cfg!(target_os = "ios")
            || cfg!(target_os = "android")
            || cfg!(target_os = "freebsd")
            || cfg!(target_os = "openbsd")
            || cfg!(target_os = "netbsd")
        {
            let res = handle_exec(
                String::from("/bin/sh"),
                vec![String::from("-c"), String::from("id -u")],
                Some(false),
            )?
            .stdout;
            let mut bool_res = false;
            if res == "1001\n" || res == "0\n" {
                bool_res = true;
            }
            assert!(bool_res);
        } else if cfg!(target_os = "macos") {
            let res = handle_exec(
                String::from("/bin/echo"),
                vec![String::from("hello")],
                Some(false),
            )?
            .stdout;
            assert_eq!(res, "hello\n");
        } else if cfg!(target_os = "windows") {
            let res = handle_exec(
                String::from("C:\\Windows\\System32\\cmd.exe"),
                vec![String::from("/c"), String::from("whoami")],
                Some(false),
            )?
            .stdout;
            let mut bool_res = false;
            if res.contains("runneradmin") || res.contains("Administrator") || res.contains("user")
            {
                bool_res = true;
            }
            assert!(bool_res);
        }
        Ok(())
    }
    #[test]
    fn test_sys_exec_complex_linux() -> anyhow::Result<()> {
        if cfg!(target_os = "linux")
            || cfg!(target_os = "ios")
            || cfg!(target_os = "macos")
            || cfg!(target_os = "android")
            || cfg!(target_os = "freebsd")
            || cfg!(target_os = "openbsd")
            || cfg!(target_os = "netbsd")
        {
            let res = handle_exec(String::from("/bin/sh"), vec![String::from("-c"), String::from("cat /etc/passwd | awk '{print $1}' | grep -E '^root:' | awk -F \":\" '{print $3}'")], Some(false))?.stdout;
            assert_eq!(res, "0\n");
        }
        Ok(())
    }

    #[test]
    fn test_sys_exec_disown_linux() -> anyhow::Result<()> {
        if cfg!(target_os = "linux")
            || cfg!(target_os = "ios")
            || cfg!(target_os = "macos")
            || cfg!(target_os = "android")
            || cfg!(target_os = "freebsd")
            || cfg!(target_os = "openbsd")
            || cfg!(target_os = "netbsd")
        {
            let tmp_file = NamedTempFile::new()?;
            let path = String::from(tmp_file.path().to_str().unwrap());
            tmp_file.close()?;

            let _res = handle_exec(
                String::from("/bin/sh"),
                vec![String::from("-c"), format!("touch {}", path.clone())],
                Some(true),
            )?;
            thread::sleep(time::Duration::from_secs(2));

            println!("{:?}", path.clone().as_str());
            assert!(Path::new(path.clone().as_str()).exists());

            let _ = fs::remove_file(path.as_str());
        }
        Ok(())
    }
    #[test]
    fn test_sys_exec_complex_windows() -> anyhow::Result<()> {
        if cfg!(target_os = "windows") {
            let res = handle_exec(
                String::from("C:\\Windows\\System32\\cmd.exe"),
                vec![
                    String::from("/c"),
                    String::from("wmic useraccount get name | findstr /i admin"),
                ],
                Some(false),
            )?
            .stdout;
            assert!(
                res.contains("runner") || res.contains("Administrator") || res.contains("user")
            );
        }
        Ok(())
    }
}

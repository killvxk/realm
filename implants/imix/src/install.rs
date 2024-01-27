use anyhow::{anyhow, Result};
use eldritch::{pb::Tome, Runtime};
use std::collections::HashMap;

pub async fn install() {
    #[cfg(debug_assertions)]
    log::info!("starting installation");

    // Iterate through all embedded files
    for embedded_file_path in eldritch::assets::Asset::iter() {
        let filename = match embedded_file_path.split(r#"/"#).last() {
            Some(local_filename) => local_filename,
            None => "",
        };

        #[cfg(debug_assertions)]
        log::debug!("checking asset {embedded_file_path}");

        // Evaluate all "main.eldritch" files
        if filename == "main.eldritch" {
            // Read eldritch content from embedded file
            #[cfg(debug_assertions)]
            log::info!("loading tome {embedded_file_path}");
            let eldritch = match load_embedded_eldritch(embedded_file_path.to_string()) {
                Ok(content) => content,
                Err(_err) => {
                    #[cfg(debug_assertions)]
                    log::error!("failed to load install asset: {_err}");

                    continue;
                }
            };

            // Run tome
            #[cfg(debug_assertions)]
            log::info!("running tome {embedded_file_path}");
            let (runtime, output) = Runtime::new();
            match tokio::task::spawn_blocking(move || {
                runtime.run(Tome {
                    eldritch,
                    parameters: HashMap::new(),
                    file_names: Vec::new(),
                });
            })
            .await
            {
                Ok(_) => {}
                Err(_err) => {
                    #[cfg(debug_assertions)]
                    log::error!("failed waiting for tome execution: {}", _err);
                }
            }

            let _output = output.collect().join("");
            #[cfg(debug_assertions)]
            log::info!("{_output}");
        }
    }
}

fn load_embedded_eldritch(path: String) -> Result<String> {
    match eldritch::assets::Asset::get(path.as_ref()) {
        Some(f) => Ok(String::from_utf8(f.data.to_vec())?),

        // {
        //     Ok(data) => data,
        //     Err(_err) => {
        //         #[cfg(debug_assertions)]
        //         log::error!("failed to load install asset: {_err}");

        //         return
        //     },
        // },
        None => {
            #[cfg(debug_assertions)]
            log::error!("no asset file at {}", path);

            return Err(anyhow!("no asset file at {}", path));
        }
    }
}

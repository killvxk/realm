import { Button } from '@chakra-ui/react';
import React from 'react';
import { useNavigate } from 'react-router-dom';
import { BarChart, Bar, Rectangle, XAxis, YAxis, CartesianGrid, Tooltip, Legend, ResponsiveContainer, Cell } from 'recharts';

import { EmptyState, EmptyStateType } from '../../../components/tavern-base-ui/EmptyState';
import { TomeTag } from '../../../utils/consts';
import { getOfflineOnlineStatus } from '../../../utils/utils';


const GroupBarChart = ({ data, loading, hosts }: { data: Array<any>, loading: boolean, hosts: Array<any> }) => {
    const navigation = useNavigate();

    if (loading) {
        return <EmptyState type={EmptyStateType.loading} label="Formatting group data..." />
    }

    const height = data.length * 40 < 320 ? 320 : data.length * 40;
    const groupWithFewestTasks = data.length > 0 ? data[0] : {};

    const getTotalActiveBeaconsForGroup = () => {
        const returnedValue = hosts.reduce((acc, curr) => {
            const matchesGroup = curr?.tags.find((tag: TomeTag) => { return tag.name === groupWithFewestTasks.name });
            const beaconStatus = getOfflineOnlineStatus(curr.beacons || [])
            if (matchesGroup) {
                return acc += beaconStatus.online;
            }
            return acc;
        }, 0);
        return returnedValue;
    };

    const activeBeaconForGroupWithFewestTasks = getTotalActiveBeaconsForGroup();

    const handleClickQuestDetails = (item: any) => {
        navigation("/tasks", {
            state: [{
                'label': item?.name,
                'kind': 'group',
                'name': item?.name,
                'value': item?.id
            }]
        })
    }

    const handleBarClick = (_: any, index?: any) => {
        const item = data[index];
        if (item.name === "undefined") {
            return null;
        }
        navigation("/tasks", {
            state: [{
                'label': item?.name,
                'kind': 'group',
                'name': item?.name,
                'value': item?.id
            }]
        })
    };

    return (
        <div className=" flex flex-col gap-6 w-full h-full">
            <div className='flex flex-row gap-4 items-center'>
                <h2 className="text-lg">Tasks by group</h2>
            </div>
            <div className='max-h-56 overflow-y-scroll'>
                <div style={{ height: `${height}px` }}>
                    <ResponsiveContainer width="100%" height="100%">
                        <BarChart
                            layout='vertical'
                            width={500}
                            height={300}
                            data={data}
                            margin={{
                                top: 5,
                                left: 5,
                                right: 5,
                                bottom: 5,
                            }}
                        >
                            <CartesianGrid strokeDasharray="3 3" />
                            <XAxis type="number" />
                            <YAxis type="category" dataKey="name" width={100} interval={0} />
                            <Tooltip />
                            <Legend />
                            <Bar dataKey="task count" fill="#553C9A" onClick={handleBarClick} activeBar={<Rectangle fill="#805AD5" stroke="#322659" />}>
                                {data.map((_, index) => (
                                    <Cell
                                        cursor="pointer"
                                        fill="#553C9A"
                                        stroke="#322659"
                                        key={`bar-cell-group-task-${index}`}
                                    />
                                ))}
                            </Bar>
                        </BarChart>
                    </ResponsiveContainer>
                </div>
            </div>
            {groupWithFewestTasks.name !== "undefined" &&
                <div className='flex flex-col border-l-4 border-purple-900 px-4 py-2 rounded'>
                    <h4 className="font-semibold text-gray-900">Consider targeting the group with fewest tasks</h4>
                    <p className='text-sm'>{groupWithFewestTasks.name} has {groupWithFewestTasks["task count"]} task run and {activeBeaconForGroupWithFewestTasks} online beacons</p>
                    <div className='flex flex-row gap-4 mt-2'>
                        <Button size="sm" variant="link" colorScheme="purple" onClick={() => {
                            handleClickQuestDetails(groupWithFewestTasks)
                        }}>
                            See quest details
                        </Button>
                    </div>
                </div>
            }
        </div>
    );
}
export default GroupBarChart;
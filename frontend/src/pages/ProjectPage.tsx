"use client";

import { useEffect, useState } from "react";
import { columns } from "@/components/task-table/Columns";
import { DataTable } from "@/components/task-table/data-table";
import { Task } from "@/lib/api/types/task";

function ProjectPage() {
    const [data, setData] = useState<Task[]>([]);
    const [loading, setLoading] = useState(true);

    useEffect(() => {
        const getData = async () => {
            const tasks: Task[] = [
                {
                    project: 42069,
                    title: "Test",
                    description: "Test",
                    deadline: "2025-05-11T14:23:45.678Z",
                    priority: "low",
                    created_by: 69,
                    assigned_to: 69,
                },
                // ...
            ];
            setData(tasks);
            setLoading(false);
        };

        getData();
    }, []);

    if (loading) return <div>Loading...</div>;

    return (

        <div className="container mx-auto py-10">
            <DataTable columns={columns} data={data} />
        </div>
    );
}

export default ProjectPage;


import { ColumnDef } from '@tanstack/react-table'
import { Task } from '@/lib/api/types/task'



export const columns: ColumnDef<Task>[] = [

    {
        accessorKey: "title",
        header: "Title",
    },

    {
        accessorKey: "priority",
        header: "Priority",
    },

    {
        accessorKey: "deadline",
        header: "Deadline",
    },

    {
        accessorKey: "assigned_to",
        header: "Assignee",
    },
]






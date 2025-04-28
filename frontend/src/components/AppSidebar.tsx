import * as React from "react"

import { SearchForm } from "@/components/SearchForm"
import { VersionSwitcher } from "@/components/VersionSwitcher"
import {
  Sidebar,
  SidebarContent,
  SidebarGroup,
  SidebarGroupContent,
  SidebarGroupLabel,
  SidebarHeader,
  SidebarMenu,
  SidebarMenuItem,
  SidebarRail,
} from "@/components/ui/sidebar"

// This is sample data.
const data = {
  versions: ["2.0"],
  navMain: [
    {
      title: "Project Overview",
      url: "#",
      items: [
        {
          title: "Project Dashboard",
          url: "#",
        },
        {
          title: "Team Members",
          url: "#",
        },
        {
          title: "Timeline",
          url: "#",
        },
        {
          title: "Tasks",
          url: "#",
        },
        {
          title: "Progress Tracker",
          url: "#",
        },
      ],
    },
    {
      title: "Task Management",
      url: "#",
      items: [
        {
          title: "Create Task",
          url: "#",
        },
        {
          title: "Task Lists",
          url: "#",
        },
        {
          title: "Task Assignments",
          url: "#",
        },
        {
          title: "Task Status",
          url: "#",
        },
        {
          title: "Due Dates",
          url: "#",
        },
        {
          title: "Dependencies",
          url: "#",
        },
      ],
    },
    {
      title: "Team Collaboration",
      url: "#",
      items: [
        {
          title: "Team Chat",
          url: "#",
        },
        {
          title: "Shared Files",
          url: "#",
        },
        {
          title: "Meeting Notes",
          url: "#",
        },
        {
          title: "Discussions",
          url: "#",
        },
        {
          title: "Announcements",
          url: "#",
        },
      ],
    },
    {
      title: "Project Settings",
      url: "#",
      items: [
        {
          title: "General Settings",
          url: "#",
        },
        {
          title: "Permissions",
          url: "#",
        },
        {
          title: "Notifications",
          url: "#",
        },
        {
          title: "Integrations",
          url: "#",
        },
        {
          title: "Billing",
          url: "#",
        },
      ],
    },
  ],
}


export function AppSidebar({ ...props }: React.ComponentProps<typeof Sidebar>) {
  return (
    <Sidebar {...props}>
      <SidebarHeader>
        <VersionSwitcher
          versions={data.versions}
          defaultVersion={data.versions[0]}
        />
        <SearchForm />
      </SidebarHeader>
      <SidebarContent>
        {/* We create a SidebarGroup for each parent. */}
        {data.navMain.map((item) => (
          <SidebarGroup key={item.title}>
            <SidebarGroupLabel>{item.title}</SidebarGroupLabel>
            <SidebarGroupContent>
              <SidebarMenu>
                {item.items.map((item) => (
                  <SidebarMenuItem key={item.title}>
                   
                  </SidebarMenuItem>
                ))}
              </SidebarMenu>
            </SidebarGroupContent>
          </SidebarGroup>
        ))}
      </SidebarContent>
      <SidebarRail />
    </Sidebar>
  )
}

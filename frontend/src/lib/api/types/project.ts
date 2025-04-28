export interface CreateProjectRequest {
    name: string;
    description: string;
    type: string;
    created_by: number;  
    attachments: any[];  
  }
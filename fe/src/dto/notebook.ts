export interface getAllNotebookResponses{
    id: string;
    name: string;
    parent_id?: string | null;
    created_at: Date;
    updated_at: Date; // Ubah dari update_at ke updated_at
    notes: GetAllNotebookResponsesNotes[] | null; // Tambahkan | null
}

export interface GetAllNotebookResponsesNotes{
   id: string;
   title: string;
   content: string;
   created_at: Date;
   updated_at: Date | null; // Ubah dari update_at ke updated_at
}

export interface CreateNotebookRequest{
    name:string,
    parent_id:string|null
}

export interface CreateNotebookResponse{
   id:string
}

export interface UpdateNotebookRequest{
   name:string;
}

export interface UpdateNotebookResponse{
   id:string;
}

export interface MoveNotebookRequest{
   parent_id:string|null;
}

export interface MoveNotebookResponse{
   id:string;
}


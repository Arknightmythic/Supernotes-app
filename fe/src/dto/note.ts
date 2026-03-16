export interface createNoteRequest{
    title: string;
    content: string;
    notebook_id: string;
}

export interface createNoteResponse{
    id: string;
}


export interface updateNoteRequest{
    title?: string;
    content?: string;
}

export interface updateNoteResponse{
    id: string;
}


export interface moveNoteRequest{
    NotebookId: string;
}

export interface moveNoteResponse{
    id: string;
}
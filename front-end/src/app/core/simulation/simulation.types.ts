export interface Simulation{
    
        id?:number,
        type: string,
        date: string,
        results: string,
        variables_entry: string,
        variables_out: string
}

export interface SearchReq{
    
    id?:number,
    type?: string,
    date?: string,
    results?: string,
    variables_entry?: string,
    variables_out?: string
}
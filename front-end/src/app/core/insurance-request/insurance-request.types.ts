export interface InsuranceRequest{
        id?: number,
        id_insurance: number,
        type: string,
        amount: number,
        justification: string,
        submission_date: string,
        response_date: string,
        status: string,
        comments: string
}

export interface SearchReq{
    id?: number,
    id_insurance?: number,
    type?: string,
    amount?: number,
    justification?: string,
    submission_date?: string,
    response_date?: string,
    status?: string,
    comments?: string
}
export interface InsurancePolicy{
        id?: number,
        client_id: number,
        insurance_id: number,
        type: string,
        amount: number,
        start_date: string,
        end_date: string,
        warranty_details: string,
        franchise: string,
        coverage_limit: number,
        policy_status: string
}

export interface SearchReq{
    id?: number,
    client_id?: number,
    insurance_id?: number,
    type?: string,
    amount?: number,
    start_date?: string,
    end_date?: string,
    warranty_details?: string,
    franchise?: string,
    coverage_limit?: number,
    policy_status?: string
}
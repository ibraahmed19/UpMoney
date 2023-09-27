export interface InsuranceProduct{
        id?: number,
        id_insurance: number,
        name: string,
        type: string,
        description: string,
        amount: number,
        coverage_limit: number,
        duration: string,
        payment_period: string,
        additional_options: string,
        status: string
}

export interface SearchReq{
    id?: number,
    id_insurance?: number,
    name?: string,
    type?: string,
    description?: string,
    amount?: number,
    coverage_limit?: number,
    duration?: string,
    payment_period?: string,
    additional_options?: string,
    status?: string
}
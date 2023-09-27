export interface Bill {
        id?: number,
        id_transaction: number,
        bill_number: string,
        client_name: string,
        iban_account: number,
        amount: number,
        issue_date: string,
        due_date: string,
        creation_date: string
      
}

export interface SearchReq{
    id?: number,
    id_transaction?: number,
    bill_number?: string,
    client_name?: string,
    iban_account?: number,
    amount?: number,
    issue_date?: string,
    due_date?: string,
    creation_date?: string
}
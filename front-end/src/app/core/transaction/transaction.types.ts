export interface Transaction{
    id?: number,
    id_banking_account: number,
    iban_sender: string,
    iban_receiver: string,
    type: TransactionType,
    status: TransactionStatus,
    amount: number,
    fee: number,
    execution_date: string,
    description: string
}

export enum TransactionStatus {
    PENDING = 'pending',
    COMPLETED = 'completed',
    ECHEC = 'echec'
}
export enum TransactionType{
    BILL, DEPOSIT
}
export interface SearchReq {
    id?: number,
    id_banking_account?: number,
    iban_sender?: string,
    iban_receiver?: string,
    type?: TransactionType,
    status?: string,
    amount?: number,
    fee?: number,
    execution_date?: string,
    description?: string
}




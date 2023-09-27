import { Component, EventEmitter, OnInit, Output, ViewEncapsulation } from '@angular/core';
import { BankingAccountService } from 'app/core/banking-account/banking-account.service';
import { Transaction, TransactionStatus, TransactionType } from 'app/core/transaction/transaction.types';
import { TransactionService } from 'app/core/transaction/transaction.service';
import { BankingAccount } from 'app/core/banking-account/banking-account.types';
import { DatePipe } from '@angular/common';

@Component({
  selector: 'app-new-transaction-form',
  templateUrl: './new-transaction-form.component.html',
  encapsulation: ViewEncapsulation.None,
})
export class NewTransactionFormComponent implements OnInit {

  userBankingAccounts: BankingAccount[]; 
  

  @Output() addTransaction = new EventEmitter<Transaction>();
  @Output() cancelForm = new EventEmitter<void>();


  transaction: Transaction = {
    id_banking_account: 1, // Initialize with a default value
    iban_sender: 'FR123456789',
    iban_receiver: '',
    type: TransactionType.BILL, // Initialize with a default value
    status: TransactionStatus.PENDING, // Initialize with a default value
    amount: 0, // Initialize with a default value
    fee: 1.5, // Initialize with a default value
    execution_date: new Date().toISOString(),
    description: '',
  };

  // Define an array of field configurations
  formFields = [
    { label: 'Type', key: 'type', type: 'select', options: ['Transfer', 'Standing Order'] },
    { label: 'Amount', key: 'amount', type: 'number' },
    { label: 'IBAN Receiver', key: 'iban_receiver', type: 'text' },
    { label: 'Description', key: 'description', type: 'text' },
  ];

  constructor(private bankingAccoountService: BankingAccountService, private transactionService: TransactionService
    ){

  }
  isFormVisible = true;

  ngOnInit(): void {
    this.bankingAccoountService.getBankingAccounts()
  }
  submitForm() {
    // Format the execution_date field
  
    this.transaction.amount = parseFloat(this.transaction.amount.toString());
    this.transactionService.newTransaction(this.transaction).subscribe((response) => {
      // Handle success or error here
      console.log(response);
    });
    this.isFormVisible = false;
  }
  

  cancelTransaction() {
    this.isFormVisible = false;
    this.cancelForm.emit();
  }
}

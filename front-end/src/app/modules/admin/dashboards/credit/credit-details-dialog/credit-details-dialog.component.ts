import { Component, Inject } from '@angular/core';
import { MAT_DIALOG_DATA } from '@angular/material/dialog';


@Component({
  selector: 'app-credit-details-dialog',
  templateUrl: './credit-details-dialog.component.html',
  styleUrls: ['./credit-details-dialog.component.scss']
})
export class CreditDetailsDialogComponent {
  constructor(@Inject(MAT_DIALOG_DATA) public data: any) { }
}

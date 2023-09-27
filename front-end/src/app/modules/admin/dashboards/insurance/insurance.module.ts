import { NgModule } from '@angular/core';
import { MatButtonModule } from '@angular/material/button';
import { MatDividerModule } from '@angular/material/divider';
import { MatIconModule } from '@angular/material/icon';
import { MatMenuModule } from '@angular/material/menu';
import { MatProgressBarModule } from '@angular/material/progress-bar';
import { MatSortModule } from '@angular/material/sort';
import { MatTableModule } from '@angular/material/table';
import { RouterModule } from '@angular/router';
import { SharedModule } from 'app/shared/shared.module';
import { NgApexchartsModule } from 'ng-apexcharts';
import { InsuranceComponent } from './insurance.component';
import { insuranceRoutes } from './insurance.routing';
import { InsuranceItemComponent } from './insurance-item/insurance-item.component';
import { MatDialogModule } from '@angular/material/dialog';
import { InsuranceDetailsComponent } from './insurance-details/insurance-details.component';

@NgModule({
    declarations: [
        InsuranceComponent,
        InsuranceItemComponent,
        InsuranceDetailsComponent
        
    ],
    imports: [
        RouterModule.forChild(insuranceRoutes),
        MatButtonModule,
        MatDividerModule,
        MatIconModule,
        MatMenuModule,
        MatProgressBarModule,
        MatSortModule,
        MatTableModule,
        NgApexchartsModule,
        SharedModule,
        MatDialogModule
    ]
})
export class InsuranceModule {
}

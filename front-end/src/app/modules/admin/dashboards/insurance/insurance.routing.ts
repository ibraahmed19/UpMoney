import { Route } from '@angular/router';
import { InsuranceComponent } from './insurance.component';
import { InsuranceResolver } from './insurance.resolvers';


export const insuranceRoutes: Route[] = [
    {
        path     : '',
        component: InsuranceComponent,
        resolve  : {
            data: InsuranceResolver
        }
    }
];

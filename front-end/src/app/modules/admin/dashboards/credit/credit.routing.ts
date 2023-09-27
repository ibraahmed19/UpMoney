import { Route } from '@angular/router';
import { CreditComponent } from './credit.component';
import { CreditResolver } from './credit.resolvers';

export const creditRoutes: Route[] = [
    {
        path     : '',
        component: CreditComponent,
        resolve  : {
            data: CreditResolver
        }
    }
];
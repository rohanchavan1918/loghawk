import { Routes } from '@angular/router';

import { DashboardComponent } from '../../dashboard/dashboard.component';
import { TagFormComponent } from '../../tag-form/tag-form.component';
import { TableListComponent } from '../../table-list/table-list.component';
import { IconsComponent } from '../../icons/icons.component';
import { LogListComponent } from '../../Log-List/Log-List.component';

export const AdminLayoutRoutes: Routes = [
    { path: 'dashboard',      component: DashboardComponent },
    { path: 'tags-form',   component: TagFormComponent },
    { path: 'tags-form/:id',   component: TagFormComponent },
    { path: 'table-list',     component: TableListComponent },
    { path: 'icons',          component: IconsComponent },
    { path: 'logs-list',        component: LogListComponent }
];

import { CommonModule } from '@angular/common';
import { Component, OnInit } from '@angular/core';
import { Router, RouterModule, RouterOutlet } from '@angular/router';
import { ButtonModule } from 'primeng/button';
import { SharedDataService } from '../../services/shared-data.service';


@Component({
    selector: 'app-categories',
    standalone: true,
    imports: [CommonModule, RouterModule, ButtonModule],
    providers: [Router, SharedDataService],
    templateUrl: './categories.component.html',
    styleUrl: './categories.component.scss'
})
export class CategoriesComponent implements OnInit {

    constructor(private svcSharedData: SharedDataService) { }

    ngOnInit(): void {
        this.svcSharedData.setBreadcrumbs('home/categories')
    }
    
}
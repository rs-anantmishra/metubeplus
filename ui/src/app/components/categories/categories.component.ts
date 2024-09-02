import { CommonModule } from '@angular/common';
import { Component, OnInit } from '@angular/core';
import { Router, RouterModule, RouterOutlet } from '@angular/router';
import { ButtonModule } from 'primeng/button';


@Component({
    selector: 'app-categories',
    standalone: true,
    imports: [CommonModule, RouterModule, ButtonModule],
    providers: [Router],
    templateUrl: './categories.component.html',
    styleUrl: './categories.component.scss'
})
export class CategoriesComponent {
    
}
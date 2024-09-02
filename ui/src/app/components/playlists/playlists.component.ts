import { CommonModule } from '@angular/common';
import { Component } from '@angular/core';
import { PaginatorModule } from 'primeng/paginator';
import { Router, RouterModule } from '@angular/router'
import { ButtonModule } from 'primeng/button';


interface PageEvent {
    first: number;
    rows: number;
    page: number;
    pageCount: number;
}

@Component({
    selector: 'app-playlists',
    standalone: true,
    imports: [CommonModule, PaginatorModule, RouterModule, ButtonModule],
    providers: [Router],
    templateUrl: './playlists.component.html',
    styleUrl: './playlists.component.scss'
})
export class PlaylistsComponent {

    visibility = 'visible'
    first: number = 0;
    rows: number = 10;

    onPageChange(event: any) {
        this.first = event.first;
        this.rows = event.rows;
    }

}
import { CommonModule } from '@angular/common';
import { Component } from '@angular/core';
import { PaginatorModule } from 'primeng/paginator';

interface PageEvent {
  first: number;
  rows: number;
  page: number;
  pageCount: number;
}

@Component({
  selector: 'app-playlists',
  standalone: true,
  imports: [CommonModule, PaginatorModule],
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
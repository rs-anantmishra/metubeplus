import { CommonModule } from '@angular/common';
import { Component, OnInit } from '@angular/core';
import { Router, RouterModule } from '@angular/router';
import { ButtonModule } from 'primeng/button';
import { SharedDataService } from '../../services/shared-data.service';

@Component({
  selector: 'app-playlist-details',
  standalone: true,
  imports: [CommonModule, RouterModule, ButtonModule],
  providers: [Router, SharedDataService],
  templateUrl: './playlist-details.component.html',
  styleUrl: './playlist-details.component.scss'
})
export class PlaylistDetailsComponent implements OnInit {
    constructor(private svcSharedData: SharedDataService) { }

    ngOnInit(): void {
        this.svcSharedData.setBreadcrumbs('home/playlists/{{playlistname}}')
    }

}

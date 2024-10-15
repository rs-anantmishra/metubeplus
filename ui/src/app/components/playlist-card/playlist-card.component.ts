import { Component, Input, OnInit } from '@angular/core';
import { PlaylistsDataResponse, PlaylistsInfo } from '../../classes/playlists'
import { CardModule } from 'primeng/card';
import { TagModule } from 'primeng/tag';
import { TooltipModule } from 'primeng/tooltip';
import { CommonModule } from '@angular/common';
import { SharedDataService } from '../../services/shared-data.service';
import { Subscription } from 'rxjs';
import { Router } from '@angular/router';

@Component({
  selector: 'app-playlist-card',
  standalone: true,
  imports: [CardModule, TagModule, TooltipModule, CommonModule],
  templateUrl: './playlist-card.component.html',
  styleUrl: './playlist-card.component.scss'
})
export class PlaylistCardComponent implements OnInit {
    @Input() playlist: PlaylistsInfo = new PlaylistsInfo();

    constructor(private sharedDataSvc: SharedDataService, private router: Router) {
    }
    
    ngOnInit(): void {
        if (this.playlist.thumbnail == '') {
            this.playlist.thumbnail = './noimage.png'
        }
    }
    
    selectedPlaylist(playlist: PlaylistsInfo) {
        console.log('clicked ', this.playlist.playlist_id)
        this.router.navigate(['/playlist-details'])
    }
}

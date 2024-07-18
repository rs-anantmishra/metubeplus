
import { Component, OnInit } from '@angular/core';
import { FormsModule } from '@angular/forms';
import { InputGroupModule } from 'primeng/inputgroup';
import { InputGroupAddonModule } from 'primeng/inputgroupaddon';
import { InputTextModule } from 'primeng/inputtext';
import { ButtonModule } from 'primeng/button';
import { CommonModule } from '@angular/common';

@Component({
  selector: 'app-downloads',
  standalone: true,
  imports: [FormsModule, InputGroupModule, InputGroupAddonModule, InputTextModule, ButtonModule, CommonModule],
  templateUrl: './downloads.component.html',
  styleUrl: './downloads.component.scss'
})
export class DownloadsComponent implements OnInit {

  placeholder = 'Video or Playlist URL'
  ngOnInit(): void {

  }

}



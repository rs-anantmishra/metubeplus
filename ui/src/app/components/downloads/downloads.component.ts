import { CommonModule } from '@angular/common';
import { Component } from '@angular/core';
import { ButtonModule } from 'primeng/button';


@Component({
  selector: 'app-downloads',
  standalone: true,
  imports: [ButtonModule, CommonModule],
  templateUrl: './downloads.component.html',
  styleUrl: './downloads.component.scss'
})
export class DownloadsComponent {

}

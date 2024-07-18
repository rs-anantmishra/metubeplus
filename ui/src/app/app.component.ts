import { CommonModule } from '@angular/common';
import { Component } from '@angular/core';
import { RouterOutlet } from '@angular/router';
import { MenuItem, MessageService } from 'primeng/api';
import { SplitButtonModule } from 'primeng/splitbutton';
import { ToastModule } from 'primeng/toast';
import { HeaderComponent } from './components/header/header.component';


@Component({
  selector: 'app-root',
  standalone: true,
  imports: [RouterOutlet, CommonModule, SplitButtonModule, ToastModule, HeaderComponent],
  templateUrl: './app.component.html',
  styleUrl: './app.component.scss',
  providers: [MessageService]
})
export class AppComponent {

  items: MenuItem[];

  constructor(private messageService: MessageService) {
    this.items = [
      {
        label: 'Update',
        command: () => {
          this.update();
        }
      },
      {
        label: 'Delete',
        command: () => {
          this.delete();
        }
      },
      { label: 'Angular Website', url: 'http://angular.io' },
      { separator: true },
      { label: 'Upload', routerLink: ['/fileupload'] }
    ];
  }

  save(severity: string) {
    this.messageService.add({ severity: severity, summary: 'Success', detail: 'Data Saved' });
  }

  update() {
    this.messageService.add({ severity: 'success', summary: 'Success', detail: 'Data Updated' });
  }

  delete() {
    this.messageService.add({ severity: 'success', summary: 'Success', detail: 'Data Deleted' });
  }
}

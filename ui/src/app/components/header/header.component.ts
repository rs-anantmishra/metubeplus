import { CommonModule } from '@angular/common';
import { Component, OnInit } from '@angular/core';
import { MenuItem, MessageService } from 'primeng/api';
import { Breadcrumb, BreadcrumbModule } from 'primeng/breadcrumb';
import { SplitButtonModule } from 'primeng/splitbutton';
import { ToastModule } from 'primeng/toast';
import { FilterService, SelectItemGroup } from 'primeng/api';
import { AutoCompleteModule } from 'primeng/autocomplete';
import { FormsModule } from '@angular/forms';
import { Router } from '@angular/router';

interface AutoCompleteCompleteEvent {
  originalEvent: Event;
  query: string;
}

@Component({
  selector: 'app-header',
  standalone: true,
  imports: [CommonModule, SplitButtonModule, ToastModule, BreadcrumbModule, AutoCompleteModule, FormsModule],
  providers: [MessageService, Router],
  templateUrl: './header.component.html',
  styleUrl: './header.component.scss'
})
export class HeaderComponent implements OnInit {


  //search-bar
  visible: string = 'hidden'

  navigationItems: MenuItem[];
  selectedCity: any;
  filteredGroups!: any[];
  groupedCities!: SelectItemGroup[];

  constructor(private router: Router, private messageService: MessageService,private filterService: FilterService) {
    this.groupedCities = [
      {
        label: 'Germany', value: 'de',
        items: [
          { label: 'Berlin', value: 'Berlin' },
          { label: 'Frankfurt', value: 'Frankfurt' },
          { label: 'Hamburg', value: 'Hamburg' },
          { label: 'Munich', value: 'Munich' }
        ]
      },
      {
        label: 'USA', value: 'us',
        items: [
          { label: 'Chicago', value: 'Chicago' },
          { label: 'Los Angeles', value: 'Los Angeles' },
          { label: 'New York', value: 'New York' },
          { label: 'San Francisco', value: 'San Francisco' }
        ]
      },
      {
        label: 'Japan', value: 'jp',
        items: [
          { label: 'Kyoto', value: 'Kyoto' },
          { label: 'Osaka', value: 'Osaka' },
          { label: 'Tokyo', value: 'Tokyo' },
          { label: 'Yokohama', value: 'Yokohama' }
        ]
      }
    ];

  //   this.items = [
  //     {
  //       label: 'Update',
  //       command: () => {
  //         this.update();
  //       }
  //     },
  //     {
  //       label: 'Delete',
  //       command: () => {
  //         this.delete();
  //       }
  //     },
  //     { label: 'Angular Website', url: 'http://angular.io' },
  //     { separator: true },
  //     { label: 'Upload', routerLink: ['/fileupload'] }
  //   ];
  // }

  this.navigationItems = [
    { label: 'Home', routerLink: ['/home'] },
    { separator: true },
    { label: 'Videos', routerLink: ['/videos'] },
    { label: 'Playlists', routerLink: ['/playlists'] },
    { label: 'Tags', routerLink: ['/tags'] },
    { label: 'Categories', routerLink: ['/categories'] },
    // { separator: true },
    // { label: 'Pattern Matching', routerLink: ['/recursive'] },
    // { label: 'Saved Patterns', routerLink: ['/notes'] },
    // { label: 'Source RegEx', routerLink: ['/source'] },
    { separator: true },
    { label: 'Activity Logs', routerLink: ['/activity-logs'] },

  ];
}

  navigateToHome() {
    this.router.navigateByUrl('/home')
  }

  update() {
    this.messageService.add({ severity: 'success', summary: 'Success', detail: 'Data Updated' });
  }

  delete() {
    this.messageService.add({ severity: 'success', summary: 'Success', detail: 'Data Deleted' });
  }

  crumbs: MenuItem[] | undefined;
  home: MenuItem | undefined;

  ngOnInit() {
    this.crumbs = [
      { label: 'Home' },
    ];

    this.home = { icon: 'pi pi-home', routerLink: '/home' };
  }

  filterGroupedCity(event: AutoCompleteCompleteEvent) {
    let query = event.query;
    let filteredGroups = [];

    for (let optgroup of this.groupedCities) {
      let filteredSubOptions = this.filterService.filter(optgroup.items, ['label'], query, "contains");
      if (filteredSubOptions && filteredSubOptions.length) {
        filteredGroups.push({
          label: optgroup.label,
          value: optgroup.value,
          items: filteredSubOptions
        });
      }
    }

    this.filteredGroups = filteredGroups;
  }

}
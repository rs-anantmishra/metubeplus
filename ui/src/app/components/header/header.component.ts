import { CommonModule, DOCUMENT } from '@angular/common';
import { Component, OnInit, inject } from '@angular/core';
import { MenuItem, MessageService } from 'primeng/api';
import { Breadcrumb, BreadcrumbModule } from 'primeng/breadcrumb';
import { SplitButtonModule } from 'primeng/splitbutton';
import { ToastModule } from 'primeng/toast';
import { FilterService, SelectItemGroup } from 'primeng/api';
import { AutoCompleteModule } from 'primeng/autocomplete';
import { FormsModule } from '@angular/forms';
import { Router, RouterModule } from '@angular/router';
import { InputSwitchModule } from 'primeng/inputswitch';

// interface AutoCompleteCompleteEvent {
//     originalEvent: Event;
//     query: string;
// }

@Component({
    selector: 'app-header',
    standalone: true,
    imports: [InputSwitchModule, CommonModule, SplitButtonModule, ToastModule, BreadcrumbModule, FormsModule],
    providers: [MessageService, Router],
    templateUrl: './header.component.html',
    styleUrl: './header.component.scss'
})
export class HeaderComponent implements OnInit {

    #document = inject(DOCUMENT);
    isDarkMode = false;
    themeIcon = ''
    activeIcon = '#dark'
    toggleTheme() {
        const linkElement = this.#document.getElementById('app-theme',) as HTMLLinkElement;
        const bodyElement = this.#document.getElementById('app-dlbg',) as HTMLBodyElement;
        if (linkElement.href.includes('light')) {
            linkElement.href = 'themes/aura-dark-blue/theme.css';
            bodyElement.className = "downloads-bg-dark"
            //this.themeIcon = 'pi pi-moon'
            this.activeIcon = '#dark'
            this.isDarkMode = true;
        } else {
            linkElement.href = 'themes/aura-light-blue/theme.css';
            bodyElement.className = "downloads-bg-light"
            //this.themeIcon = 'pi pi-sun'
            this.activeIcon = '#light'
            this.isDarkMode = false;
        }
    }


    //search-bar
    visible: string = 'hidden'

    navigationItems: MenuItem[];
    selectedCity: any;
    filteredGroups!: any[];
    groupedCities!: SelectItemGroup[];

    constructor(private router: Router, private messageService: MessageService, private filterService: FilterService) {
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

    navigateToHome(e: any) {
        this.router.navigate(['/home']);
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

    // filterGroupedCity(event: AutoCompleteCompleteEvent) {
    //     let query = event.query;
    //     let filteredGroups = [];

    //     for (let optgroup of this.groupedCities) {
    //         let filteredSubOptions = this.filterService.filter(optgroup.items, ['label'], query, "contains");
    //         if (filteredSubOptions && filteredSubOptions.length) {
    //             filteredGroups.push({
    //                 label: optgroup.label,
    //                 value: optgroup.value,
    //                 items: filteredSubOptions
    //             });
    //         }
    //     }

    //     this.filteredGroups = filteredGroups;
    // }

}
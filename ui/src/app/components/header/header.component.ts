import { CommonModule, DOCUMENT } from '@angular/common';
import { HostListener, Directive, Component, OnInit, inject } from '@angular/core';
import { MenuItem, MessageService } from 'primeng/api';
import { Breadcrumb, BreadcrumbModule } from 'primeng/breadcrumb';
import { SplitButtonModule } from 'primeng/splitbutton';
import { ToastModule } from 'primeng/toast';
import { FilterService, SelectItemGroup } from 'primeng/api';
import { AutoCompleteCompleteEvent, AutoCompleteModule } from 'primeng/autocomplete';
import { FormsModule } from '@angular/forms';
import { Router, RouterModule } from '@angular/router';
import { InputSwitchModule } from 'primeng/inputswitch';
import { SharedDataService } from '../../services/shared-data.service';
import { BehaviorSubject } from 'rxjs';

// interface AutoCompleteCompleteEvent {
//     originalEvent: Event;
//     query: string;
// }

@Component({
    selector: 'app-header',
    standalone: true,
    imports: [InputSwitchModule, CommonModule, SplitButtonModule, ToastModule, FormsModule, AutoCompleteModule],
    providers: [MessageService, Router, SharedDataService],
    templateUrl: './header.component.html',
    styleUrl: './header.component.scss'
})
export class HeaderComponent implements OnInit {

    #document = inject(DOCUMENT);
    themeIcon = ''
    activeIcon = '#dark'
    toggleTheme() {
        const linkElement = this.#document.getElementById('app-theme',) as HTMLLinkElement;
        const bodyElement = this.#document.getElementById('app-dlbg',) as HTMLBodyElement;

        if (linkElement.href.includes('light')) {
            this.setDarkMode();
        } else {
            this.setLightMode();
        }
    }

    setLightMode() {
        const linkElement = this.#document.getElementById('app-theme',) as HTMLLinkElement;
        const bodyElement = this.#document.getElementById('app-dlbg',) as HTMLBodyElement;

        linkElement.href = 'themes/aura-light-blue/theme.css';
        bodyElement.className = "downloads-bg-light"
        this.activeIcon = '#dark'
        this.sharedDataSvc.setIsDarkMode(false)
    }

    setDarkMode() {
        const linkElement = this.#document.getElementById('app-theme',) as HTMLLinkElement;
        const bodyElement = this.#document.getElementById('app-dlbg',) as HTMLBodyElement;

        linkElement.href = 'themes/aura-dark-blue/theme.css';
        bodyElement.className = "downloads-bg-dark"
        this.activeIcon = '#light'
        this.sharedDataSvc.setIsDarkMode(true)
    }

    //search-bar
    visible: string = 'visible'

    navItems: MenuItem[];
    selectedCity: any;
    filteredGroups!: any[];
    groupedCities!: SelectItemGroup[];

    constructor(private router: Router, private messageService: MessageService, private filterService: FilterService, private sharedDataSvc: SharedDataService) {

        //check and set theme
        let isDarkMode = this.sharedDataSvc.getIsDarkMode();
        if (isDarkMode === null) {
            this.sharedDataSvc.setIsDarkMode(true)
        } else {
            if (isDarkMode === true) {
                this.setDarkMode()
            } else if (isDarkMode === false) {
                this.setLightMode()
            }
        }

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

        this.navItems = [
            { label: 'Home', routerLink: ['/home'], command: () => { this.navigate('/home'); } },
            { separator: true },
            { label: 'Videos', routerLink: ['/videos'], command: () => { this.navigate('/videos'); } },
            { label: 'Playlists', routerLink: ['/playlists'], command: () => { this.navigate('/playlists'); } },
            { label: 'Tags', routerLink: ['/tags'], command: () => { this.navigate('/tags'); } },
            { label: 'Categories', routerLink: ['/categories'], command: () => { this.navigate('/categories'); } },
            // { separator: true },
            // { label: 'Pattern Matching', routerLink: ['/recursive'] },
            // { label: 'Saved Patterns', routerLink: ['/notes'] },
            // { label: 'Source RegEx', routerLink: ['/source'] },
            { separator: true },
            { label: 'Activity Logs', routerLink: ['/activity-logs'], command: () => { this.navigate('/logs'); } },
        ];
    }

    navigate(route: string) {
        this.router.navigate([route]);
    }

    crumbs: MenuItem[] | undefined;
    home: MenuItem | undefined;
    crumbsSubscription!: BehaviorSubject<string>;

    ngOnInit() {

    }

    //combinations Alt + Shift + P = Playlists
    //combinations Alt + Shift + V = Videos
    //combinations Alt + Shift + C = Channels
    //combinations Alt + Shift + L = Logs
    @HostListener("document:keydown", ["$event"]) handleKeyboardEvent(event: KeyboardEvent) {
        if (event.key === 'P' && event.altKey) {
            this.navigate('/playlists')
        }
        if (event.key === 'V' && event.altKey) {
            this.navigate('/videos')
        }
        if (event.key === 'C' && event.altKey) {
            this.navigate('/channels')
        }
        if (event.key === 'L' && event.altKey) {
            this.navigate('/logs')
        }
        if ((event.key === 'H' || event.key === 'D') && event.altKey) {
            this.navigate('/home')
        }
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
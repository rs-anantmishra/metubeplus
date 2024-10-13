import { CommonModule, DOCUMENT } from '@angular/common';
import { HostListener, Directive, Component, OnInit, inject, OnDestroy, effect } from '@angular/core';
import { MenuItem, MessageService } from 'primeng/api';
import { Breadcrumb, BreadcrumbModule } from 'primeng/breadcrumb';
import { SplitButtonModule } from 'primeng/splitbutton';
import { ToastModule } from 'primeng/toast';
import { FilterService, SelectItemGroup } from 'primeng/api';
import { AutoCompleteCompleteEvent, AutoCompleteModule } from 'primeng/autocomplete';
import { FormsModule } from '@angular/forms';
import { ActivatedRoute, Router, RouterModule } from '@angular/router';
import { InputSwitchModule } from 'primeng/inputswitch';
import { SharedDataService } from '../../services/shared-data.service';
import { VideosService } from '../../services/videos.service';
import { BehaviorSubject, Subscription } from 'rxjs';
import { ContentSearchResponse, ContentSearch } from '../../classes/search';
import { group } from '@angular/animations';

// interface AutoCompleteCompleteEvent {
//     originalEvent: Event;
//     query: string;
// }

@Component({
    selector: 'app-header',
    standalone: true,
    imports: [InputSwitchModule, CommonModule, SplitButtonModule, ToastModule, FormsModule, AutoCompleteModule],
    providers: [MessageService, Router, SharedDataService, VideosService],
    templateUrl: './header.component.html',
    styleUrl: './header.component.scss'
})
export class HeaderComponent implements OnInit, OnDestroy {

    isHomepage = true
    isHomepageSub!: Subscription;

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
    filteredGroups!: any[];
    selectedTitle: any;
    groupedTitles!: SelectItemGroup[];

    constructor(private router: Router,
        private messageService: MessageService,
        private videosSvc: VideosService,
        private filterService: FilterService,
        private sharedDataSvc: SharedDataService) {


        //isHomepage
        this.isHomepageSub = this.sharedDataSvc.getIsHomepage().subscribe(() => { this.setIsHomepage() })

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
    ngOnDestroy(): void {
        //unsubscribe
        this.isHomepageSub.unsubscribe()
    }

    navigate(route: string) {
        this.router.navigate([route]);
    }

    crumbs: MenuItem[] | undefined;
    home: MenuItem | undefined;
    crumbsSubscription!: BehaviorSubject<string>;

    async ngOnInit(): Promise<void> {
        let result = await this.videosSvc.getContentSearchInfo()
        this.groupedTitles = await this.buildAutoCompleteDataset(result)
    }

    async buildAutoCompleteDataset(raw: ContentSearchResponse): Promise<SelectItemGroup[]> {
        let content: ContentSearch[] = raw.data;
        let result: SelectItemGroup[] = [];

        //group titles by channel
        let grouped = content.reduce(
            (result: any, currentValue: any) => {
                (result[currentValue['channel']] = result[currentValue['channel']] || []).push({ "label": currentValue['title'], 'value': currentValue['video_id'] });
                return result;
            }, {});

        //format json in requires manner
        for (let key in grouped) {
            if (grouped.hasOwnProperty(key)) {
                let val = { label: key, value: "", items: grouped[key] };
                result.push(val)
            }
        }
        return result;
    }

    //keyboard shortcuts
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

    filterGroupedContent(event: AutoCompleteCompleteEvent) {
        let query = event.query;
        let filteredGroups = [];

        for (let optgroup of this.groupedTitles) {
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

    setIsHomepage() {
        console.log('called from header')
        this.isHomepage = false;
        console.log(`The current value is: ${this.sharedDataSvc.getIsPageHome()}`);
    }
}
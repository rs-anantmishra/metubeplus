import { Routes } from '@angular/router';
import { DownloadsComponent } from './components/downloads/downloads.component';
import { VideosComponent } from './components/videos/videos.component';
import { TagsComponent } from './components/tags/tags.component';
import { CategoriesComponent } from './components/categories/categories.component';
import { VideoDetailsComponent } from './components/video-details/video-details.component';
import { PlaylistsComponent } from './components/playlists/playlists.component';
import { PlaylistDetailsComponent } from './components/playlist-details/playlist-details.component';

export const routes: Routes = [
    { path: '', redirectTo: '/home', pathMatch: 'full' },
    { path: 'home', component: DownloadsComponent, title: 'Metube+' },
    { path: 'videos', component: VideosComponent, title: 'Videos+' },
    { path: 'tags', component: TagsComponent, title: 'Tags+' },
    { path: 'categories', component: CategoriesComponent, title: 'Categories+' },
    { path: 'play', component: VideoDetailsComponent, title: 'Metube+' },
    { path: 'playlists', component: PlaylistsComponent, title: 'Playlists+' },
    { path: 'playlist-details', component: PlaylistDetailsComponent, title: 'Metube+' },
    { path: '**', redirectTo: '/home'}
];

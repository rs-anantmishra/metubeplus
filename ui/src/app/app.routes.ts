import { Routes } from '@angular/router';
import { DownloadsComponent } from './components/downloads/downloads.component';
import { VideosComponent } from './components/videos/videos.component';
import { PlaylistsComponent } from './components/playlists/playlists.component';
import { CategoriesComponent } from './components/categories/categories.component';
import { TagsComponent } from './components/tags/tags.component';
import { PlaylistDetailsComponent } from './components/playlist-details/playlist-details.component';
import { VideoDetailsComponent } from './components/video-details/video-details.component';

export const routes: Routes = [
    { path: 'home', component: DownloadsComponent, title: 'Metube+' },
    { path: 'videos', component: VideosComponent, title: 'Videos+' },
    { path: 'playlists', component: PlaylistsComponent, title: 'Playlists+' },
    { path: 'tags', component: TagsComponent, title: 'Tags+' },
    { path: 'categories', component: CategoriesComponent, title: 'Categories+' },
    { path: 'video-details', component: VideoDetailsComponent, title: 'Metube+' },
    { path: 'playlist-details', component: PlaylistDetailsComponent, title: 'Metube+' },
    { path: '', redirectTo: '/home', pathMatch: 'full' },
    { path: '**', redirectTo: 'home'}
];

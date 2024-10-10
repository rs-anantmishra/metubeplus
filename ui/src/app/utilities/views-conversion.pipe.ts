// views-conversion.pipe.ts
import { Pipe, PipeTransform } from '@angular/core';
@Pipe({
  name: 'minifiedViewCount',
  standalone: true,
})
export class MinifiedViewCount implements PipeTransform {
  transform(bytes: number, decimals = 2): string {
    if (!+bytes) return '0 views'

    const k = 1000
    const dm = decimals < 0 ? 0 : decimals
    const sizes = ['', 'K views', 'M views', 'B views', 'T views']

    const i = Math.floor(Math.log(bytes) / Math.log(k))
    let result = `${parseFloat((bytes / Math.pow(k, i)).toFixed(dm))}${sizes[i]}`

    return result
  }
}
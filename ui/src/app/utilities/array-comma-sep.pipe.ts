// array-comma-sep.pipe.ts
import { Pipe, PipeTransform } from '@angular/core';
@Pipe({
    name: 'commaSepStringFromArray',
    standalone: true,
})
export class CommaSepStringFromArray implements PipeTransform {
    transform(value: string[]): string {
        if (value.length > 0) {
            value = value.map((line) => `#${line}`);
            let result = value.join(', ')
            return result
        } else {
            return 'Video does not have any tags.'
        }
    }
}
import {format, parse} from 'date-fns';
import { ru } from 'date-fns/locale';

export function datetimeToBString(datetime, localeString="ru") {
    return format(datetime, 'PPP HH:mm:ss', {locale: ru})
}

export function datetimeToString(datetime, localeString="ru") {
    return format(datetime, 'dd-MM-yyyy HH:mm', {locale: ru})
}

export function stringToDate(datetime, localeString="ru") {
    return parse(datetime, 'PPP HH:mm:ss', new Date(), {locale: ru})
}
import * as en from './en.json';
import * as fr from './fr.json';

export function getTranslations(locale: string): {[key: string]: string} {
    switch (locale) {
    case 'en':
        return en;
    case 'fr':
        return fr;
    }
    return {};
}


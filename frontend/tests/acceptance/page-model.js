import { Selector, t } from 'testcafe';

export default class Page {
    constructor() {
        this.view = Selector('div.p-view-select', {timeout: 15000});
        this.camera = Selector('div.p-camera-select', {timeout: 15000});
        this.countries = Selector('div.p-countries-select', {timeout: 15000});
        this.time = Selector('div.p-time-select', {timeout: 15000});
        this.search1 = Selector('div.p-search-field input', {timeout: 15000});
    }

    async setFilter(filter, option) {

        const SelectOption = await Selector('a').withText(option);
        switch (filter) {
            case 'view':
                await t
                    .click(this.view, {timeout: 15000});
                break;
            case 'camera':
                await t
                    .click(this.camera, {timeout: 15000});
                break;
            case 'time':
                await t
                    .click(this.time, {timeout: 15000});
                break;
            case 'countries':
                await t
                    .click(this.countries, {timeout: 15000});
                break;
            default:
        }
        await t
            .click(Selector('a').withText(option), {timeout: 15000} )
    }

    async search(term) {
        await t
            .typeText(this.search1, term)
            .pressKey('enter')
    }

    async openNav() {
        if (await Selector('button.p-navigation-show').visible) {
            await t.click(Selector('button.p-navigation-show'));
        } else if (await Selector('div.p-navigation-expand').exists) {
            await t.click(Selector('div.p-navigation-expand i'));
        }
    }
}

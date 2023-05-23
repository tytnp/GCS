
import {readdirSync} from 'fs'
import {resolveModule} from 'local-pkg'
import {dirname} from 'path'

let iconPkgMap;

const iconPkgs = [
    '@vicons/fluent',
    '@vicons/ionicons4',
    '@vicons/ionicons5',
    '@vicons/antd',
    '@vicons/fa',
    '@vicons/material',
    '@vicons/tabler',
    '@vicons/carbon',
];

export function ViconsResolver() {
    if (!iconPkgMap) {
        try {
            iconPkgMap = new Map();
            iconPkgs.forEach((pkg) => {
                const icons = readdirSync(dirname(resolveModule(pkg)), { withFileTypes: true })
                    .filter((item) => !item.isDirectory() && item.name.match(/^[A-Z][A-Za-z0-9]+\.js$/))
                    .map((item) => item.name.replace(/\.js$/, ''));
                icons.forEach((icon) => iconPkgMap.set(icon, pkg));
            });
        } catch (error) {
            console.error(error);
            throw new Error('[unplugin-vue-components] 加载 vicons 失败，请确保已安装。');
        }
    }
    return {
        type: 'component',
        resolve: (name) => {
            if (iconPkgMap.has(name)) {
                let as;
                if (name === 'Map') {
                    as = iconPkgMap.get(name).split('/')[1] + 'Map';
                }
                return {
                    as: as,
                    name: name,
                    from: iconPkgMap.get(name),
                };
            }
        },
    };
}
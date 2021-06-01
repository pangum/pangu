module.exports = ctx => ({
    dest: './dist',
    locales: {
        '/': {
            lang: 'zh-CN',
            title: '盘古',
            description: 'Golang应用快速开发框架'
        }
    },
    head: [
        ['link', {rel: 'icon', href: `/logo.png`}],
        ['link', {rel: 'manifest', href: '/manifest.json'}],
        ['link', {rel: "icon", type: "image/png", sizes: "32x32", href: "/icons/favicon-32x32.png"}],
        ['link', {rel: "icon", type: "image/png", sizes: "16x16", href: "/icons/favicon-16x16.png"}],
        ['meta', {name: 'theme-color', content: '#3eaf7c'}],
        ['meta', {name: 'apple-mobile-web-app-capable', content: 'yes'}],
        ['meta', {name: 'apple-mobile-web-app-status-bar-style', content: 'black'}],
        ['link', {rel: 'apple-touch-icon', href: `/icons/pangu-152x152.png`}],
        ['link', {rel: 'mask-icon', href: '/icons/safari-pinned-tab.svg', color: '#3eaf7c'}],
        ['meta', {name: 'msapplication-TileImage', content: '/icons/pangu-144x144.png'}],
        ['meta', {name: 'msapplication-TileColor', content: '#000000'}]
    ],
    markdown: {
        lineNumbers: true
    },
    themeConfig: {
        repo: 'storezhang/pangu',
        editLinks: true,
        docsDir: 'doc/docs',
        algolia: ctx.isProd ? ({
            apiKey: 'f7edee00640ed06f44542bb62b4d9e5b',
            indexName: 'doc',
            algoliaOptions: {
                facetFilters: ['tags:v1']
            }
        }) : null,
        smoothScroll: true,
        locales: {
            '/': {
                label: '简体中文',
                selectText: '选择语言',
                ariaLabel: '选择语言',
                editLinkText: '在GitHub上编辑此页',
                lastUpdated: '上次更新',
                nav: require('./nav/zh'),
                sidebar: {
                    '/api/': getApiSidebar(),
                    '/guide/': getGuideSidebar('指南', '深入'),
                    '/config/': getConfigSidebar('系统', '命令行'),
                    '/plugin/': getPluginSidebar('插件', '官方插件'),
                }
            }
        }
    },
    plugins: [
        ['@vuepress/back-to-top', true],
        ['@vuepress/pwa', {
            serviceWorker: true,
            updatePopup: true
        }],
        ['@vuepress/medium-zoom', true],
        ['@vuepress/google-analytics', {
            ga: 'UA-128189152-1'
        }],
        ['container', {
            type: 'vue',
            before: '<pre class="vue-container"><code>',
            after: '</code></pre>'
        }],
        ['container', {
            type: 'upgrade',
            before: info => `<UpgradePath title="${info}">`,
            after: '</UpgradePath>'
        }],
        ['flowchart']
    ],
    extraWatchFiles: [
        '.vuepress/nav/en.js',
        '.vuepress/nav/zh.js'
    ]
})

function getApiSidebar() {
    return [
        'cli',
        'node'
    ]
}

function getGuideSidebar(groupA, groupB) {
    return [{
        title: groupA,
        collapsable: false,
        children: [
            '',
            'start',
            'concept',
            'config',
            'migration',
        ]
    }, {
        title: groupB,
        collapsable: false,
        children: [
            'di',
            'serve',
            'command',
            'arg',
            'version',
            'plugin'
        ]
    }]
}

function getConfigSidebar(groupA, groupB) {
    return [{
        title: groupA,
        collapsable: false,
        children: [
            '',
            'banner',
            'default',
            'validator'
        ]
    }, {
        title: groupB,
        collapsable: false,
        children: [
            'name',
            'usage',
            'description',
            'authors',
            'copyright'
        ]
    }]
}

function getPluginSidebar(groupA, groupB) {
    return [{
        title: groupA,
        collapsable: false,
        children: [
            '',
            'using',
            'writing'
        ]
    }, {
        title: groupB,
        collapsable: false,
        children: [
            'database'
        ]
    }]
}

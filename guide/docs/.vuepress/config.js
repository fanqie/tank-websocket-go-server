module.exports = {
  base: '/tank-websocket-go-server/',
  title: 'Tank WebSocket',
  description: 'A lightweight, feature-rich WebSocket server implementation in Go',
  routerMode: 'hash',
  themeConfig: {
    logo: '/images/logo.png',
    nav: [
      { text: 'Home', link: '/en/' },
      { text: 'Guide', link: '/en/guide/' }
    ],
    sidebar: {
      '/en/guide/': [
        {
          title: 'Guide',
          collapsable: false,
          children: [
            '',
            'installation',
            'quick-start',
            'client-connection',
            'heartbeat',
            'topic-subscription',
            'authentication',
            'debug-logging'
          ]
        }
      ]
    }
  },
  locales: {
    '/': {
      lang: 'en-US',
      title: 'Tank WebSocket',
      description: 'A lightweight, feature-rich WebSocket server implementation in Go',
      selectText: 'Languages',
      label: 'English',
      ariaLabel: 'Languages',
      nav: [
        { text: 'Home', link: '/en/' },
        { text: 'Guide', link: '/en/guide/' }
      ],
      sidebar: {
        '/en/guide/': [
          {
            title: 'Guide',
            collapsable: false,
            children: [
              '',
              'installation',
              'quick-start',
              'client-connection',
              'heartbeat',
              'topic-subscription',
              'authentication',
              'debug-logging'
            ]
          }
        ]
      }
    },
    '/zh/': {
      lang: 'zh-CN',
      title: 'Tank WebSocket',
      description: '一个用 Go 语言实现的轻量级、功能丰富的 WebSocket 服务器',
      selectText: '选择语言',
      label: '简体中文',
      ariaLabel: '语言',
      nav: [
        { text: '首页', link: '/zh/' },
        { text: '指南', link: '/zh/guide/' }
      ],
      sidebar: {
        '/zh/guide/': [
          {
            title: '指南',
            collapsable: false,
            children: [
              '',
              'installation',
              'quick-start',
              'client-connection',
              'heartbeat',
              'topic-subscription',
              'authentication',
              'debug-logging'
            ]
          }
        ]
      }
    }
  },
  plugins: [
    '@vuepress/back-to-top',
    '@vuepress/medium-zoom'
  ],
  head: [
    ['link', { rel: 'icon', href: '/images/favicon.ico' }]
  ],
  markdown: {
    lineNumbers: true
  },
  configureWebpack: {
    output: {
      publicPath: '/tank-websocket-go-server/'
    }
  }
} 
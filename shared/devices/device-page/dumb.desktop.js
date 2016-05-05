// @flow
import Render from './index.render'
import type {DumbComponentMap} from '../../constants/types/more'

const common = {
  type: 'desktop',
  name: 'Home Computer',
  currentDevice: false,
  revokedAt: false,
  onRevoke: () => {}
}

const map: DumbComponentMap<Render> = {
  component: Render,
  mocks: {
    'Normal': {
      ...common,
      deviceID: 'aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa',
      currentDevice: true,
      timeline: [
        {type: 'LastUsed',
          desc: 'Last used yesterday'},
        {type: 'Added',
          desc: 'Added Mar 03, 2015'}
      ],
      banner: {
        type: 'OutOfDate',
        desc: 'Home Computer is running an outdated version of Keybase. Remember to update!'
      }
    },
    'Revoked': {
      ...common,
      deviceID: 'bbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbb',
      currentDevice: false,
      revokedAt: 1444423192000,
      timeline: [
        {type: 'Revoked',
          desc: 'Revoked yesterday',
          subDesc: 'Home Computer'},
        {type: 'LastUsed',
          desc: 'Last used Nov 12, 2015',
          subDesc: '83 days ago'},
        {type: 'Added',
          desc: 'Added Mar 03, 2014',
          subDesc: 'Home Computer'}
      ]
    },
    'Unlock': {
      ...common,
      deviceID: 'cccccccccccccccccccccccccccccccc',
      name: 'Chris\'s iPhone',
      type: 'mobile',
      timeline: [
        {type: 'LastUsed',
          desc: 'Last used Mar 25, 2016',
          subDesc: '16 days ago'},
        {type: 'Added',
          desc: 'Added Mar 03, 2015',
          subDesc: 'Home Computer'}
      ],
      banner: {
        type: 'WillUnlock',
        desc: 'Turning on this device will unlock 6 of your private folders.'
      }
    },
    'Paper': {
      ...common,
      deviceID: 'dddddddddddddddddddddddddddddddd',
      name: 'project green...',
      type: 'backup',
      timeline: [
        {type: 'Added',
          desc: 'Created Mar 03, 2014',
          subDesc: 'Home Computer'}
      ]
    }
  }
}

export default {
  'Device Page': map
}

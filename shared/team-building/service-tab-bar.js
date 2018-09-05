// @flow
import React from 'react'
import * as Kb from '../common-adapters/index'
import * as Styles from '../styles'
import {serviceIdToLogo16} from './shared'
import type {ServiceIdWithContact} from '../constants/team-building'

// TODO
// * Add styles for mobile
// * Add contact icon
// * Add tooltip
// * Add highlighted underline

const SERVICE_LIST = ['keybase', 'contact', 'twitter', 'facebook', 'github', 'reddit', 'hackernews']

type Props = {
  selectedService: ServiceIdWithContact,
  onChangeService: (newService: ServiceIdWithContact) => void,
  serviceResultCount: {[key: ServiceIdWithContact]: ?number},
  showServiceResultCount: boolean,
}

type IconProps = {
  service: ServiceIdWithContact,
  onClick: () => void,
  count: ?number,
  showCount: boolean,
  isActive: boolean,
}

const ServiceIcon = (props: IconProps) => (
  <Kb.ClickableBox onClick={props.onClick}>
    <Kb.Box2 direction="horizontal" centerChildren={true} style={styles.serviceIconContainer}>
      <Kb.Icon
        type={serviceIdToLogo16(props.service, props.isActive)}
        style={Styles.collapseStyles([
          styles.serviceIcon,
          styles[props.isActive ? 'activeIcon' : 'inactiveIcon'],
        ])}
      />
      {!!props.showCount &&
        (props.count ? (
          <Kb.Text type="tiny-semibold" style={styles.resultCount}>
            {props.count}
          </Kb.Text>
        ) : (
          <Kb.Icon type="icon-progress-grey-animated" style={styles.progressIcon} />
        ))}
    </Kb.Box2>
  </Kb.ClickableBox>
)

const ServiceTabBar = (props: Props) => (
  <Kb.Box2 direction="horizontal">
    {SERVICE_LIST.map(service => (
      <ServiceIcon
        key={service}
        service={service}
        onClick={() => props.onChangeService(service)}
        count={props.showServiceResultCount[service]}
        showCount={props.showServiceResultCount}
        isActive={props.selectedService === service}
      />
    ))}
  </Kb.Box2>
)

const styles = Styles.styleSheetCreate({
  container: Styles.platformStyles({
    isElectron: {
      height: 40,
      marginLeft: Styles.globalMargins.small,
      marginTop: Styles.globalMargins.large, // small
      width: 370,
    },
    common: {
      ...Styles.globalStyles.rounded,
      borderColor: Styles.globalColors.black_20,
      borderWidth: 1,
      borderStyle: 'solid',
    },
  }),
  serviceIconContainer: {
    marginLeft: Styles.globalMargins.xtiny,
    marginRight: Styles.globalMargins.xtiny,
  },
  serviceIcon: {},
  activeIcon: {},
  inactiveIcon: {},
  progressIcon: {
    height: 10,
    width: 10,
  },
  resultCount: {},
})

export default ServiceTabBar

// @flow
import * as React from 'react'
import {Box2, Text, Icon, Input, WaitingButton} from '../../../common-adapters'
import {globalStyles, globalColors, globalMargins, platformStyles, styleSheetCreate} from '../../../styles'
import QRImage from './qr-image'
import QRScan from './qr-scan'
import type {Tab, DeviceType} from './helper'

type Props = {
  currentDeviceAlreadyProvisioned: boolean,
  currentDeviceType: DeviceType,
  otherDeviceType: DeviceType,
  defaultMode: Tab,
  validModes: Array<Tab>,
  enterQrCodeInstructions: string,
  enterTextCodeInputHint: string,
  enterTextCodeInstructions: string,
  isValidLookingCode: string => boolean,
  onSubmitTextCode: (textCode: string) => void,
  viewQrCode: string,
  viewQrCodeInstructions: string,
  viewTextCode: string,
  viewTextCodeInstructions: string,
}

type State = {
  mode: Tab,
}

const PanelContainer = ({instructions, children}) => (
  <Box2 direction="vertical" style={styles.panelContainer} gap="medium" gapStart={true}>
    <Text type="Body" style={styles.panelInstructions}>
      {instructions}
    </Text>
    {children}
  </Box2>
)

const ViewQR = ({url, instructions}) => (
  <PanelContainer instructions={instructions}>
    <QRImage url={url} />
  </PanelContainer>
)

const ScanQR = ({onScan, instructions}) => (
  <PanelContainer instructions={instructions}>
    <QRScan onScan={onScan} />
  </PanelContainer>
)

const ViewText = ({code, instructions}) => (
  <PanelContainer instructions={instructions}>
    <Text type="Terminal" style={styles.textCode}>
      {code}
    </Text>
  </PanelContainer>
)

class EnterText extends React.Component<
  {
    isValidLookingCode: string => boolean,
    onSubmit: string => void,
    instructions: string,
    inputHint: string,
  },
  {value: string, canSubmit: boolean}
> {
  state = {canSubmit: false, value: ''}
  _onSubmit = () => this.props.onSubmit(this.state.value)
  _updateValue = value => this.setState({canSubmit: this.props.isValidLookingCode(value), value})
  render() {
    return (
      <PanelContainer instructions={this.props.instructions}>
        <Input
          uncontrolled={true}
          onEnterKeyDown={this._onSubmit}
          onChangeText={this._updateValue}
          hintText={this.props.inputHint}
        />
        <WaitingButton
          type="Primary"
          label="Continue"
          onClick={this._onSubmit}
          disabled={!this.state.canSubmit}
          waitingKey="TODO"
        />
      </PanelContainer>
    )
  }
}

class CodePage2 extends React.Component<Props, State> {
  constructor(props: Props) {
    super(props)
    this.state = {
      mode: this.props.defaultMode,
    }
  }

  componentDidUpdate(prevProps: Props) {
    if (this.props.defaultMode !== prevProps.defaultMode) {
      this.setState({
        mode: this.props.defaultMode,
      })
    }
  }

  _tabsMap = {
    enterText: 'Enter a text code',
    scanQR: 'Scan a QR Code',
    viewQR: 'See a QR Code',
    viewText: 'See a text code',
  }

  _getContent = () => {
    switch (this.state.mode) {
      case 'scanQR':
        return (
          <ScanQR onScan={this.props.onSubmitTextCode} instructions={this.props.viewQrCodeInstructions} />
        )
      case 'viewQR':
        return <ViewQR url={this.props.viewQrCode} instructions={this.props.viewQrCodeInstructions} />
      case 'enterText':
        return (
          <EnterText
            isValidLookingCode={this.props.isValidLookingCode}
            instructions={this.props.enterTextCodeInstructions}
            inputHint={this.props.enterTextCodeInputHint}
            onSubmit={this.props.onSubmitTextCode}
          />
        )
      case 'viewText':
        return <ViewText code={this.props.viewTextCode} instructions={this.props.viewTextCodeInstructions} />
      default:
        return null
    }
  }

  render() {
    return (
      <Box2 direction="vertical">
        <Box2 direction="horizontal" gap="small">
          {this.props.validModes.map(mode => (
            <Text
              key={mode}
              type={mode === this.state.mode ? 'BodySecondaryLink' : 'BodyPrimaryLink'}
              style={styles.tab}
              onClick={() => this.setState({mode})}
            >
              {this._tabsMap[mode]}
            </Text>
          ))}
        </Box2>
        {this._getContent()}
      </Box2>
    )
  }
}

const styles = styleSheetCreate({
  panelContainer: {
    alignItems: 'center',
  },
  panelInstructions: {
    maxWidth: 300,
    textAlign: 'center',
  },
  tabs: {
    padding: globalMargins.small,
  },
  textCode: {
    borderRadius: 4,
    borderStyle: 'solid',
    borderWidth: 1,
    color: globalColors.darkBlue,
    maxWidth: 300,
    padding: 20,
    textAlign: 'center',
  },
})

export default CodePage2
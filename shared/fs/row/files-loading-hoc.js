// @flow
import * as I from 'immutable'
import * as React from 'react'
import {compose, namedConnect} from '../../util/container'
import * as FsGen from '../../actions/fs-gen'
import * as Types from '../../constants/types/fs'

const mapStateToProps = state => ({
  syncingPaths: state.fs.uploads.syncingPaths,
})

const mapDispatchToProps = (dispatch, {path}) => ({
  loadFolderList: () => dispatch(FsGen.createFolderListLoad({path, refreshTag: 'main'})),
  loadFavorites: () => dispatch(FsGen.createFavoritesLoad()),
})

const mergeProps = ({syncingPaths}, {loadFolderList, loadFavorites}, o) => ({
  syncingPaths,
  loadFolderList,
  loadFavorites,
  ...o,
})

type FilesLoadingHocProps = {
  syncingPaths: I.Set<Types.Path>,
  loadFolderList: () => void,
  loadFavorites: () => void,
  path: Types.Path,
}

const FilesLoadingHoc = (ComposedComponent: React.ComponentType<any>) =>
  class extends React.PureComponent<FilesLoadingHocProps> {
    _load = () => {
      const pathLevel = Types.getPathLevel(this.props.path)
      if (pathLevel < 2) {
        return
      }
      pathLevel === 2 && this.props.loadFavorites()
      // This is needed not only inside in a tlf, but also in tlf list, to get
      // `writable` for tlf root.
      this.props.loadFolderList()
    }
    componentDidMount() {
      this._load()
    }
    componentDidUpdate(prevProps) {
      // This gets called on route changes too, e.g. when user clicks the
      // action menu. So only load folder list when path changes.
      this.props.path !== prevProps.path && this._load()
    }
    render() {
      return <ComposedComponent {...this.props} />
    }
  }

export default compose(
  namedConnect(mapStateToProps, mapDispatchToProps, mergeProps, 'ConnectedFilesLoadingHoc'),
  FilesLoadingHoc
)

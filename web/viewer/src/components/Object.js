import SoundBoard from './SoundBoard';

const Object = ({object, players}) => {
  if(!object) {
    return null;
  }

  switch (object.type) {
    case 'obstacle':
      return <div className="obstacle" />
    case 'prop':
      switch (object.id) {
        case 0:
            return <div className="prop bullet" />
          break;
        case 1:
            return <div className="prop bomb"><SoundBoard.Shot2 /></div>
          break;
        default:
          return <div className="prop" >{object.id}</div>
      }
    case 'event':
      switch (object.id) {
        case 3:
            return <div className="event shot-hit"><SoundBoard.Shot3 /></div>
          break;
        case 5:
            return <div className="event bomb-explode"><SoundBoard.Explosion1 /></div>
          break;
        case 6:
            return <div className="event player-kill" ><SoundBoard.Shot3 /><SoundBoard.Scream1 /></div>
          break;
        default:
          return null;
      }
    case 'player':
    let styles = {};
    if(players) {
      styles = players[object.id].styles;
    }

      return (
        <div className={`player player-${object.id}`}>
          <div style={{cssText:styles.head}} className="player-head"></div>
          <div style={{cssText:styles.body}} className="player-body">
            <div style={{cssText:styles.arm}} className="player-arm"></div>
            <div style={{cssText:styles.arm}} className="player-arm"></div>
          </div>
          <div style={{cssText:styles.feet}} className="player-feet">
            <div style={{cssText:styles.foot}} className="player-foot"></div>
            <div style={{cssText:styles.foot}} className="player-foot"></div>
          </div>
          <SoundBoard.Move1 />
        </div>
      );
    default:
      return null;
  }
}

export default Object;

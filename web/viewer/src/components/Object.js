const Object = ({object}) => {
  if(!object) {
    return null;
  }
  switch (object.type) {
    case 'obstacle':
      return <div className="obstacle" />
    case 'prop':
      switch (object.id) {
        case 1:
            return <div className="prop bullet" />
          break;
        case 2:
            return <div className="prop bomb" />
          break;
        default:
          return <div className="prop" >{object.id}</div>
      }
    case 'event':
      switch (object.id) {
        case 4:
            return <div className="event bomb-explode" />
          break;
        default:
          return <div className="event" >{object.id}</div>
      }
    case 'player':
      return (
        <div className={`player player-${object.id}`}>
          <div className="player-head" />
          <div className="player-body" />
          <div className="player-feet" />
        </div>
      );
    default:
      return null;
  }
}

export default Object;

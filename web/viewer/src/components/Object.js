const Object = ({object}) => {
  if(!object) {
    return null;
  }
  switch (object.type) {
    case 'obstacle':
      return <div className="obstacle" />
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

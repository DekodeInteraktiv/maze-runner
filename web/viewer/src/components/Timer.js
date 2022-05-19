const Timer = ({timer, active, roundTime}) => {
  if(!timer || !active) {
    return null;
  } else {
    let time = roundTime - timer;
    const minutes = parseInt(time / 60);
    return (<div className="timer">{minutes}:{time - minutes * 60}</div>);
  }
}

export default Timer;

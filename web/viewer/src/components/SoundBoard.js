import select from '../audio/select.wav';
import blip1 from '../audio/blip1.wav';
import blip2 from '../audio/blip2.wav';
import shot1 from '../audio/shot1.wav';
import shot2 from '../audio/shot2.wav';
import shot3 from '../audio/shot3.wav';
import move1 from '../audio/move1.wav';
import explosion1 from '../audio/explosion1.wav';
import scream1 from '../audio/scream1.wav';

const SoundBoard = {
  Select: () => {
    return <audio src={select} autoPlay />
  },
  Blip1: () => {
    return <audio src={blip1} autoPlay />
  },
  Blip2: () => {
    return <audio src={blip2} autoPlay />
  },
  Shot1: () => {
    return <audio src={shot1} autoPlay />
  },
  Shot2: () => {
    return <audio src={shot2} autoPlay />
  },
  Shot3: () => {
    return <audio src={shot3} autoPlay />
  },
  Move1: () => {
    return <audio src={move1} autoPlay />
  },
  Explosion1: () => {
    return <audio src={explosion1} autoPlay />
  },
  Scream1: () => {
    return <audio src={scream1} autoPlay />
  }
}

export default SoundBoard;

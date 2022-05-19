const data = {
	map: [
		[
			{color: 1, object: false},
			{color: 1, object: false},
			{color: 1, object: false},
			{color: 2, object: {type: 'player', id: 2}},
			{color: false, object: false},
			{color: false, object: {type: 'obstacle'}},
			{color: false, object: false},
			{color: false, object: false},
			{color: false, object: false},
			{color: false, object: false}
		],
		[
			{color: false, object: false},
			{color: false, object: false},
			{color: false, object: false},
			{color: false, object: false},
			{color: false, object: {type: 'obstacle'}},
			{color: false, object: {type: 'obstacle'}},
			{color: false, object: {type: 'obstacle'}},
			{color: 2, object: false},
			{color: false, object: false},
			{color: false, object: false}
		],
		[
			{color: 3, object: {type: 'player', id: 3}},
			{color: false, object: false},
			{color: false, object: false},
			{color: false, object: false},
			{color: false, object: false},
			{color: false, object: false},
			{color: false, object: false},
			{color: false, object: false},
			{color: false, object: false},
			{color: false, object: false}
		],
		[
			{color: false, object: false},
			{color: false, object: false},
			{color: false, object: false},
			{color: false, object: false},
			{color: false, object: false},
			{color: false, object: false},
			{color: false, object: false},
			{color: false, object: false},
			{color: false, object: false},
			{color: false, object: {type: 'obstacle'}}
		],
		[
			{color: false, object: false},
			{color: false, object: false},
			{color: false, object: false},
			{color: false, object: {type: 'obstacle'}},
			{color: false, object: {type: 'obstacle'}},
			{color: false, object: {type: 'obstacle'}},
			{color: 3, object: false},
			{color: false, object: false},
			{color: 3, object: false},
			{color: false, object: {type: 'obstacle'}}
		],
		[
			{color: false, object: false},
			{color: false, object: false},
			{color: false, object: false},
			{color: false, object: false},
			{color: false, object: {type: 'obstacle'}},
			{color: false, object: {type: 'obstacle'}},
			{color: false, object: false},
			{color: false, object: false},
			{color: false, object: false},
			{color: false, object: false}
		],
		[
			{color: false, object: false},
			{color: false, object: false},
			{color: false, object: false},
			{color: 4, object: false},
			{color: 4, object: false},
			{color: 4, object: false},
			{color: false, object: false},
			{color: false, object: false},
			{color: false, object: false},
			{color: false, object: false}
		],
		[
			{color: false, object: false},
			{color: false, object: {type: 'obstacle'}},
			{color: false, object: {type: 'obstacle'}},
			{color: false, object: false},
			{color: false, object: false},
			{color: false, object: false},
			{color: false, object: false},
			{color: false, object: false},
			{color: 2, object: {type: 'player', id: 1}},
			{color: 2, object: false}
		],
		[
			{color: false, object: false},
			{color: false, object: {type: 'obstacle'}},
			{color: false, object: {type: 'obstacle'}},
			{color: false, object: false},
			{color: false, object: false},
			{color: false, object: false},
			{color: false, object: false},
			{color: false, object: false},
			{color: false, object: false},
			{color: 2, object: false}
		],
		[
			{color: false, object: false},
			{color: false, object: false},
			{color: false, object: false},
			{color: false, object: false},
			{color: false, object: false},
			{color: false, object: false},
			{color: false, object: {type: 'obstacle'}},
			{color: false, object: {type: 'obstacle'}},
			{color: false, object: false},
			{color: 2, object: false}
		],
	],
	players: [
		0,
		{
			name: 'Cats',
			color: 'red',
		},
		{
			name: 'Cows',
			color: 'blue'
		},
		{
			name: 'Monkeys',
			color: 'green'
		},
		{
			name: 'Pigs',
			color: 'pink'
		}
	],
	meta: {
		columns: 12,
		row: 30,
		state: 'active',
		timer: 20000
	}
};

export default data;

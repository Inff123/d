![1](https://github.com/ectrc/snow/assets/13946988/64c3b1ac-d308-4e5d-ad8d-2b7aead29195)

# Snow

Performance first, universal Fortnite backend written in Go.

## Features

- **Blazing Fast** Written in Go, snow is extremely fast and can handle thousands of requests per second with its built-in caching system.
- **Profile Changes** Snow keeps track of profile changes exactly like Fortnite does, meaning it is one-to-one with the game.

## Examples

```golang
user := person.NewPerson()
snapshot := user.AthenaProfile.Snapshot()

quest := person.NewQuest("Quest:Quest_1", "ChallengeBundle:Daily_1", "ChallengeBundleSchedule:Paid_1")
{
  quest.AddObjective("quest_objective_eliminateplayers", 0)
  quest.AddObjective("quest_objective_top1", 0)
  quest.AddObjective("quest_objective_place_top10", 0)

  quest.UpdateObjectiveCount("quest_objective_eliminateplayers", 10)
  quest.UpdateObjectiveCount("quest_objective_place_top10", -3)

  quest.RemoveObjective("quest_objective_top1")
}
user.AthenaProfile.Quests.AddQuest(quest)

giftBox := person.NewGift("GiftBox:GB_Default", 1, user.ID, "Hello, Bully!")
{
  giftBox.AddLoot(person.NewItemWithType("AthenaCharacter:CID_002_Athena_Commando_F_Default", 1, "athena"))
}
user.CommonCoreProfile.Gifts.AddGift(giftBox)

currency := person.NewItem("Currency:MtxPurchased", 100)
user.CommonCoreProfile.Items.AddItem(currency)

user.Save()
user.AthenaProfile.Diff(snapshot)
```

## What's next?

- Be able to convert my person structures into the required format for the game. This mainly targets the profiles and their changes.
- Person Authentication for the game to determine if the person is valid or not. Fortnite uses JWT tokens for this which makes it easy to implement.
- Embed game assets into the backend e.g. Game XP Curve, Quest Data etc. _This would mean a single binary that can be run anywhere without the need of external files._
- Implement the HTTP API for the game to communicate with the backend. This is the most important part of the project as it needs to handle thousands of requests per second. _Should I use Fiber?_
- Interact with external Buckets to save player data externally.
- A way to interact with persons outside of the game. This is mainly for a web app and other services to interact with the backend.
- Game Server Communication. This would mean a websocket server that communicates with the game servers to send and receive data.

## Contributing

Contributions are welcome! Please open an issue or pull request if you would like to contribute.

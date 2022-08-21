# GOAL
Go Auto Layout

https://developer.apple.com/library/archive/documentation/UserExperience/Conceptual/AutolayoutPG/index.html

Try and implement a CLI version of auto layout in go.

Goal (pun intended) is to try and learn auto layout in a deep way.

# Example Screenshot

![screenshot1][screenshot1]{height=200px width=433px}

[//]:
[screenshot1]: ./screenshot1.png "screenshot1"

# Example Json

```
{
  "root": {
    "subviews": [
      {
        "name": "view1",
        "class": "UIView",
        "subviews": [
          {
            "name": "view2",
            "class": "UIView",
            "subviews": [
              {
                "name": "label1",
                "class": "UILabel",
                "text": "Hello"
              }
            ]
          },
          {
            "name": "view3",
            "class": "UIView",
            "subviews": [
              {
                "name": "label2",
                "class": "UILabel",
                "text": "There"
              }
            ]
          }
        ]
      }
    ]
  }
}
```

# Example Swift

```
        let view1 = UIView()
        view1.backgroundColor = .cyan
        view1.translatesAutoresizingMaskIntoConstraints = false
        view.addSubview(view1)
        
        let view2 = UIView()
        view2.backgroundColor = .orange
        view2.translatesAutoresizingMaskIntoConstraints = false
        view1.addSubview(view2)
        
        let view3 = UIView()
        view3.backgroundColor = .red
        view3.translatesAutoresizingMaskIntoConstraints = false
        view1.addSubview(view3)
        
        let label1 = UILabel()
        label1.text = "Hello"
        label1.textColor = .black
        label1.translatesAutoresizingMaskIntoConstraints = false
        view2.addSubview(label1)
        
        let label2 = UILabel()
        label2.text = "There"
        label2.textColor = .white
        label2.translatesAutoresizingMaskIntoConstraints = false
        view3.addSubview(label2)

        NSLayoutConstraint.activate([
            view1.leadingAnchor.constraint(equalTo: view.leadingAnchor, constant: 33),
            view1.topAnchor.constraint(equalTo: view.topAnchor, constant: 33),
            view1.trailingAnchor.constraint(equalTo: view.trailingAnchor, constant: -33),
            view1.bottomAnchor.constraint(equalTo: view.bottomAnchor, constant: -33),
             ])

        NSLayoutConstraint.activate([
            label1.leadingAnchor.constraint(equalTo: view2.leadingAnchor, constant: 33),
            label1.topAnchor.constraint(equalTo: view2.topAnchor, constant: 33),
             ])

        NSLayoutConstraint.activate([
            view2.leadingAnchor.constraint(equalTo: view1.leadingAnchor, constant: 33),
            view2.topAnchor.constraint(equalTo: view1.topAnchor, constant: 33),
            view2.trailingAnchor.constraint(equalTo: label1.trailingAnchor, constant: 33),
            view2.bottomAnchor.constraint(equalTo: view1.bottomAnchor, constant: -33),
             ])
        
        NSLayoutConstraint.activate([
            label2.leadingAnchor.constraint(equalTo: view3.leadingAnchor, constant: 33),
            label2.topAnchor.constraint(equalTo: view3.topAnchor, constant: 33),
             ])

        NSLayoutConstraint.activate([
            view3.leadingAnchor.constraint(equalTo: label1.leadingAnchor, constant: 99),
            view3.topAnchor.constraint(equalTo: view1.topAnchor, constant: 33),
            view3.trailingAnchor.constraint(equalTo: label2.trailingAnchor, constant: 33),
            view3.bottomAnchor.constraint(equalTo: view1.bottomAnchor, constant: -33),
             ])
```

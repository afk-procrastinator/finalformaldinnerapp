//
//  SecondViewController.swift
//  formaldinner
//
//  Created by Spencer Michaels on 2/18/20.
//  Copyright © 2020 cate. All rights reserved.
//

import UIKit

class SecondViewController: UIViewController, UITextFieldDelegate {
    
    // Outlet declarations:
    @IBOutlet weak var TableLabel: UILabel!
    @IBOutlet weak var SearchBar: UITextField!
    @IBOutlet weak var TableOrService: UILabel!
    @IBOutlet weak var First: UILabel!
    @IBOutlet weak var Second: UILabel!
    @IBOutlet weak var Third: UILabel!
    @IBOutlet weak var Fourth: UILabel!
    @IBOutlet weak var Fifth: UILabel!
    @IBOutlet weak var Sixth: UILabel!
    @IBOutlet weak var Seventh: UILabel!
    @IBOutlet weak var WaiterLabel: UILabel!
    
    @IBOutlet weak var WaiterIs: UILabel!
    @IBOutlet weak var seatedWith: UILabel!
    
    // Variable declarations:
    var idValue = [Int:Int]()
    var seating = [SeatingModel]()
    var allNames = [String]()
    
    
    override func viewDidLoad() {
        
        super.viewDidLoad()
        
        // Sets delegate to UITextFieldDelegate
        SearchBar?.delegate = self
        
        // Loads in the JSON
        loadJSON()
        
        print("testing123")
    }
    
    // Memory warning catch:
    override func didReceiveMemoryWarning() {
        super.didReceiveMemoryWarning()
    }
    
    
    // Loads in and parses the JSON from my server:
    func loadJSON() {
        let jsonUrlString = "http://localhost:6969/"
        guard let url = URL(string:jsonUrlString) else {return}
        URLSession.shared.dataTask(with: url) { (data, response, err) in
            guard let data = data else { return }
            do {
                self.seating = try! JSONDecoder().decode([SeatingModel].self, from: data)
            }
        } .resume()
    }
    
    // Takes a person with position int from seating, then prints their table to the label
    func pullID(int: Int, seating: [SeatingModel], label: UILabel) {
        print("IDSEAT:",seating[int].Table)
        let IsKC = checkIfNotSeated(int: int).0
        let IsW = checkIfNotSeated(int: int).1
        if IsKC == false && IsW == false {
            TableOrService.text = "Seated at:"
            label.text = String(seating[int].Table)
        } else if IsKC == true {
            TableOrService.text = "Kitchen Crew!"
            eliminateOtherTexts()
            label.text = ""
        } else if IsW == true {
            TableOrService.text = "Waiter at:"
            //WaiterLabel.text = "You!"
            label.text = String(seating[int].Value-10)
            
        }
        
    }
    
    // Finds everyone in the JSON with the same table number, then prints it. Need to implement a way to keep it from including waiters and KC.
    func pullTablemates(int: Int, seating: [SeatingModel]) {
        var tableMatesAr = [String]()
        print("pulled")
        let ogPerson = seating[int].Table
        for seat in seating {
            if seat.Table == ogPerson && seat.Value != seating[int].Value {
                let fullName = seat.First + " " + seat.Last
                tableMatesAr.append(fullName)
            }
        }
        let waiterPerson = ogPerson + 10
        WaiterLabel.text = seating[waiterPerson].First + " " + seating[waiterPerson].Last
        First.text = tableMatesAr[0]
        Second.text = tableMatesAr[1]
        Third.text = tableMatesAr[2]
        Fourth.text = tableMatesAr[3]
        Fifth.text = tableMatesAr[4]
        Sixth.text = tableMatesAr[5]
        Seventh.text = tableMatesAr[6]
        
    }
    
    // Function to check if the person is KC or W. First bool is True if they're a KC member, second is True if they are a W.
    func checkIfNotSeated(int:Int) -> (Bool, Bool) {
        if int >= 0 && int <= 9 {
            print("W")
            return(true, false)
        } else if int >= 10 && int <= 41 {
            print("KC")
            eliminateOtherTexts()
            return(false, true)
        } else {
            print("neither")
            return (false, false)
        }
        
    }
    
    func eliminateOtherTexts() {
        First.text = ""
        Second.text = ""
        Third.text = ""
        Fourth.text = ""
        Fifth.text = ""
        Sixth.text = ""
        Seventh.text = ""
        WaiterLabel.text = ""
        seatedWith?.text = ""
    }
    
    // ---------------------- SEARCHBAR FUNCTIONS: ----------------------
    
    // Magnifying glass activates search:
    @IBAction func searchPressed(_ sender: UIButton) {
        SearchBar.endEditing(true)
        print("pressed")
    }
    
    // Enter activates search:
    func textFieldShouldReturn(_ textField: UITextField) -> Bool {
        SearchBar.endEditing(true)
        return true
    }
    
    // Resets the filler text:
    func textFieldShouldEndEditing(_ textField: UITextField) -> Bool {
        if textField.text != "" {
            return true
        } else {
            textField.placeholder = "Type something"
            return false
        }
    }
    
    // When editing ends:
    func textFieldDidEndEditing(_ textField: UITextField) {
        // Pulls the name from the searchbar:
        let nameToSearch = SearchBar.text
        // Loops through the main array, checks names against the nameToSearch
        for seats in seating {
            if nameToSearch == seats.First + " " + seats.Last {
                // Approve match:
                print("match: \(seats.First + seats.Last)")
                pullTablemates(int: seats.Value, seating: seating)

                pullID(int: seats.Value, seating: seating, label: TableLabel)
            }
        }
        SearchBar.text = ""
    }
    
    
    
    
}




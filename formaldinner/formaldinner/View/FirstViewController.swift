//
//  FirstViewController.swift
//  formaldinner
//
//  Created by Spencer Michaels on 3/2/20.
//  Copyright © 2020 cate. All rights reserved.
//

import UIKit

class FirstViewController: UIViewController  {
    
    
   // @IBOutlet weak var testLabel: UILabel!
    @IBOutlet weak var testLabel: UILabel!
    
    var passTest:String = ""
    
    override func viewDidLoad() {
        super.viewDidLoad()
        
        testLabel?.text = passTest
        // Do any additional setup after loading the view.
    }
    

}

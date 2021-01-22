#include <iostream>

using namespace std;

struct node{
    int data;
    node *next;
};

class Stack{
    node *top;

public:
    Stack(): top(NULL){}
    void push(int);
    void pop();
    void display();
};

void Stack::push(int value){
    node *new_node = new node;
    new_node->data = value;

    if(top == NULL){
        new_node->next = NULL;
        top = new_node;
    }else{
        new_node->next = top;
        top = new_node;
    }
}

void Stack::pop(){
    if(top == NULL){
        cout << "The stack is empty !\n";
    }else{
        node *ptr = top;
        top = top->next;

        cout << "The poped node is: "<<ptr->data<<endl;

        delete ptr;
    }
}

void Stack::display(){
    if(top == NULL){
        cout << "The Stack is empty !\n";
    }else{
        node *ptr = top;
        while(ptr != NULL){
            cout<<ptr->data<<" ";
            ptr = ptr->next;
        }
        cout<<endl;
    }
}

int main(){
    int option, value;
    Stack s;

    do
    {
        cout << "\n**** MENU ****\n"
             << "1. Push\n"
             << "2. Pop\n"
             << "3. Display\n"
             << "4. Exit\n";

        cout << "Enter your option: ";
        cin >> option;

        switch (option)
        {
        case 1:
            cout << "Enter the value to be pushed:";
            cin >> value;

            s.push(value);
            break;
        case 2:
            s.pop();
            break;
        case 3:
            s.display();
            break;
        case 4:
            break;
        default:
            cout << "Wrong option !\n";
            break;
        }
    } while (option != 4);

    if (option == 4)
    {
        cout << "\nThank you for using the program !\n"
             << "Have a good day. \n\n";
    }
}
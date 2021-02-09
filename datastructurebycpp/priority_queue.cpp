#include<iostream>
using namespace std;

struct node{
    int data;
    int priority;
    node *next;
};

class Queue{
    node *start;
    public:
        Queue():start(NULL){}
        void insert(node *);
        void del(node *);
        void display(node *);
        void setStart(node *);
        node *getStart();
};

node *Queue::getStart(){
    return start;
}

void Queue::setStart(node *start){
    this->start = start;
}

void Queue::insert(node *start){
    int value, priority;

    cout<<"Enter the value to be inserted: ";
    cin>>value;

    cout<<"Enter the priority: ";
    cin>>priority;

    node *new_node = new node;
    new_node->data = value;
    new_node->priority = priority;

    if(start == NULL || priority < start->priority){
        new_node->next = start;
        start = new_node;
    }else{
        node *ptr = start;
        while(ptr->next != NULL && ptr->next->priority <= priority){
            ptr = ptr->next;
        }
        new_node->next = ptr->next;
        ptr->next = new_node;
    }
    this->start = start;
}

void Queue::del(node *start){
    if(start == NULL){
        cout<<"The Queue is empty !\n";
    }else{
        node *ptr = start;
        cout<<"The deleted node is: "<<ptr->data<<endl;
        start = start->next;
        delete ptr;
    }
    this->start = start;
}

void Queue::display(node *start){
    if(start == NULL){
        cout<<"The Queue is empty !\n";
    }else{
        node *ptr = start;
        while(ptr != NULL){
            cout<<ptr->data<<" [Priority: "<<ptr->priority<<"]"<<endl;
            ptr = ptr->next;
        }
    }
}

int main(){
    int option;
    Queue q;
    q.setStart(NULL);

    do{
        cout<<"\n************ MENU *************\n"
            <<"1. Insert\n"
            <<"2. Delete\n"
            <<"3. Display\n"
            <<"4. Exit\n";
        
        cout<<"Enter your option: ";
        cin>>option;

        switch(option){
            case 1:
                q.insert(q.getStart());
                break;
            case 2:
                q.del(q.getStart());
                break;
            case 3:
                q.display(q.getStart());
                break;
            case 4:
                break;
            default:
                cout<<"Wrong option !\n";
                break;
        }
    }while(option != 4);

    if(option==4){
        cout<<"\nThank you for using the program !\n"
            <<"Have a good day !\n";
    }
}
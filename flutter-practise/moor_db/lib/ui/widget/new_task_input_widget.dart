import 'package:flutter/material.dart';
import 'package:moor_db/data/moor_database.dart';
import 'package:provider/provider.dart';

class NewTaskInput extends StatefulWidget {
  NewTaskInput({Key key}) : super(key: key);

  @override
  _NewTaskInputState createState() {
    return _NewTaskInputState();
  }
}

class _NewTaskInputState extends State<NewTaskInput> {
  DateTime newTaskDate;
  TextEditingController controller;
  @override
  void initState() {
    super.initState();
    controller = TextEditingController();
  }

  @override
  void dispose() {
    super.dispose();
  }

  @override
  Widget build(BuildContext context) {
    // TODO: implement build
    return Container(
      padding: const EdgeInsets.all(8.0),
      child: Row(
        mainAxisSize: MainAxisSize.max,
        children: [
          _buildTextField(context),
          _buildDateButton(context),
        ],
      ),
    );
  }

  Expanded _buildTextField(BuildContext context){
    return Expanded(
      child: TextField(
        controller: controller,
        decoration: InputDecoration(hintText: 'Task name'),
        onSubmitted: (inputName){
          final database = Provider.of<AppDatabase>(context);
          final task = Task(
            name: inputName,
            dueData: newTaskDate,
          );
          database.insertTask(task);
          resetValuesAfterSubmit();
        },
      ),
    );
  }

  IconButton _buildDateButton(BuildContext context){
    return IconButton(
      icon: Icon(Icons.calendar_today),
      onPressed: () async{
        newTaskDate = await showDatePicker(
          context: context,
          initialDate: DateTime.now(),
          firstDate: DateTime(2010),
          lastDate: DateTime(2050),
        );
      },
    );
  }

  void resetValuesAfterSubmit(){
    setState(() {
      newTaskDate = null;
      controller.clear();
    });
  }
}
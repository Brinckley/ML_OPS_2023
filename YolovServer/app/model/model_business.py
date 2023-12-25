# # -*- coding: utf-8 -*-
# """MLSD(2).ipynb
#
# Automatically generated by Colaboratory.
#
# Original file is located at
#     https://colab.research.google.com/drive/1TinjrBJf5_OY9JqU3rHJdMP8UfL-j_ub
# """
#
# !pip3 install ultralytics
#
# !pip3 install onnx
#
# !pip3 install jax
#
# from ultralytics import YOLO
#
# from ultralytics import YOLO
#
# # Load a model
# model1 = YOLO('./yolov8n-cls.pt')  # load an official model
# model2 = YOLO('./yolov8n-seg.pt')
#
# # Predict with the model
# # results = model1('https://ultralytics.com/images/bus.jpg', save = True)  # predict on an image
# results = model2('https://ultralytics.com/images/bus.jpg', save = True)  # predict on an image
#
# !yolo export model='./yolov8n-seg.pt' format=engine device=0
#
# from ultralytics import YOLO
#
# # Load a model
# model2 = YOLO('./yolov8n-seg.engine')
#
# # Predict with the model
# # results = model1('https://ultralytics.com/images/bus.jpg', save = True)  # predict on an image
# results = model2('https://ultralytics.com/images/bus.jpg', save = True)  # predict on an image
#
# print(results)
#
# import torch
#
# # Download COCO val
# torch.hub.download_url_to_file('https://ultralytics.com/assets/coco2017val.zip', 'tmp.zip')
# !unzip -q tmp.zip -d ../datasets && rm tmp.zip  # unzip
#
# import locale
# locale.getpreferredencoding = lambda: "UTF-8"
#
# # Download COCO val
# import torch
# torch.hub.download_url_to_file('https://ultralytics.com/assets/coco2017val.zip', 'tmp.zip')  # download (780M - 5000 images)
# !unzip -q tmp.zip -d datasets && rm tmp.zip  # unzip
#
# from ultralytics.utils.benchmarks import benchmark
#
# benchmark(model='yolov8n-seg.pt', imgsz=640)
#
# !wget http://images.cocodataset.org/zips/train2014.zip
#
# !unzip train2014.zip